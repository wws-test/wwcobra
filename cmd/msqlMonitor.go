package cmd

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/color"
)

var (
	// 颜色定义
	Yellow     = color.Yellow.Render
	Cyan       = color.Cyan.Render
	LightGreen = color.Style{color.Green, color.OpBold}.Render
)

var (
	logfile string
	db      *sql.DB
	cstZone = time.FixedZone("CST", 8*3600)
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "实时输出 MySQL 监控日志",
	Long:  `'monitor' 子命令将打印指定 SQL 语句`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initDB(); err != nil {
			// 如果初始化过程中出现错误，记录错误日志并退出程序
			log.Fatalf("initDB error: %s", err)
		}

		// 延迟执行函数，用于清理资源
		defer func() {
			// 尝试清理日志文件，如果发生错误，记录错误日志
			if err := cleanGenerakLog(); err != nil {
				log.Printf("cleanGenerakLog error: %s \n", err)
			}
			// 尝试关闭日志记录，如果发生错误，记录错误日志
			if err := closeLogRaw(); err != nil {
				log.Printf("closeLogRaw error: %s \n", err)
			}
			// 尝试关闭数据库连接，如果发生错误，记录错误日志
			if err := db.Close(); err != nil {
				log.Printf("close database connection error: %s \n", err)
			}

			// 打印程序退出信息
			fmt.Println("\nBye hacker :)")
		}()

		// 输出程序启动信息
		fmt.Println("start mysql monitor...")

		// 设置MySQL日志输出
		if err := setMySQLLogOutput(); err != nil {
			// 如果设置过程中出现错误，记录错误日志并退出程序
			log.Fatalf("setMySQLLogOutput error: %s", err)
		}

		// 打开日志记录的原始模式
		if err := openLogRaw(); err != nil {
			// 如果打开过程中出现错误，记录错误日志并退出程序
			log.Fatalf("openLogRaw error: %s", err)
		}
		// Check if the log file exists
		if _, err := os.Stat(logfile); os.IsNotExist(err) {
			fmt.Printf("Log file does not exist: %s\n", logfile)
			prompt := promptui.Prompt{
				Label:   "请输入mysql日志文件的路径",
				Default: "/opt/metersphere/data/mysql/04c6a6eceb9a.log",
				Validate: func(input string) error {
					if strings.TrimSpace(input) == "" {
						return fmt.Errorf("文件路径不能为空")
					}
					return nil
				},
			}
			result, err := prompt.Run()
			if err != nil {
				fmt.Printf("提示失败 %v\n", err)
				return
			}
			logfile = result
			// Optionally, you can validate the new path here
			fmt.Printf("Using new log file path: %s\n", logfile)
		}

		// 在这里添加监控逻辑
		watchdog()
	},
}

func init() {
	// 启用前缀匹配
	cobra.EnablePrefixMatching = true
	rootCmd.AddCommand(monitorCmd)
}
func main() {
	// 在非Windows环境下，检查当前用户是否为root
	if runtime.GOOS != "windows" && !isRoot() {
		log.Fatalln("run as a user with root! Thx:)")
	}

	// 初始化数据库连接
	if err := initDB(); err != nil {
		// 如果初始化过程中出现错误，记录错误日志并退出程序
		log.Fatalf("initDB error: %s", err)
	}

	// 延迟执行函数，用于清理资源
	defer func() {
		// 尝试清理日志文件，如果发生错误，记录错误日志
		if err := cleanGenerakLog(); err != nil {
			log.Printf("cleanGenerakLog error: %s \n", err)
		}
		// 尝试关闭日志记录，如果发生错误，记录错误日志
		if err := closeLogRaw(); err != nil {
			log.Printf("closeLogRaw error: %s \n", err)
		}
		// 尝试关闭数据库连接，如果发生错误，记录错误日志
		if err := db.Close(); err != nil {
			log.Printf("close database connection error: %s \n", err)
		}

		// 打印程序退出信息
		fmt.Println("\nBye hacker :)")
	}()

	// 输出程序启动信息
	fmt.Println("start mysql monitor...")

	// 设置MySQL日志输出
	if err := setMySQLLogOutput(); err != nil {
		// 如果设置过程中出现错误，记录错误日志并退出程序
		log.Fatalf("setMySQLLogOutput error: %s", err)
	}

	// 打开日志记录的原始模式
	if err := openLogRaw(); err != nil {
		// 如果打开过程中出现错误，记录错误日志并退出程序
		log.Fatalf("openLogRaw error: %s", err)
	}

	// 启动看门狗程序，用于监控日志文件
	watchdog()
}

// initDB 函数用于初始化数据库连接
func initDB() error {
	var err error
	// 使用配置的参数连接MySQL数据库
	db, err = sql.Open("mysql", fmt.Sprintf("root:Password123@mysql@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"))
	// 如果连接失败，返回错误
	if err != nil {
		return err
	}

	// 设置数据库的最大打开连接数
	db.SetMaxOpenConns(20)
	// 设置数据库的最大空闲连接数
	db.SetMaxIdleConns(10)
	// 成功初始化，返回nil
	return nil
}

// catMySQLVersion 函数用于获取MySQL数据库的版本信息
func catMySQLVersion() (string, error) {
	var version string                         // 定义字符串变量用于存储MySQL版本信息
	row := db.QueryRow("SELECT version();")    // 执行SQL查询以获取MySQL版本
	if err := row.Scan(&version); err != nil { // 尝试将查询结果扫描到version变量中
		return "", err // 如果扫描过程中发生错误，返回错误信息
	}
	return version, nil // 成功获取版本信息后，返回版本字符串和nil表示无错误
}

// 打开日志记录的原始模式
func openLogRaw() error {
	version, err := catMySQLVersion()
	// 获取MySQL版本信息，若有错误则返回
	if err != nil {
		return err
	}
	vs := strings.Split(version, ".")
	// 将版本字符串按点分割，若分割结果不足一项则返回错误
	if len(vs) < 1 {
		return fmt.Errorf("mysql version '%s' ", version)
	}

	if v, err := strconv.Atoi(vs[0]); err != nil {
		// 将版本号的第一部分转换为整数，若转换失败则返回错误
		return err
	} else if v < 8 {
		// 若MySQL版本号小于8，则无需设置log_raw，直接返回nil
		return nil
	}
	// sett log_raw=1
	if _, err := db.Exec("SET GLOBAL log_raw = 'ON'"); err != nil {
		// 尝试将log_raw设置为'ON'，若执行失败则返回错误
		return err
	}
	// 成功设置log_raw，返回nil
	return nil
}

func closeLogRaw() error {
	// sett log_raw=0
	if _, err := db.Exec("SET GLOBAL log_raw = 'OFF'"); err != nil {
		return err
	}
	return nil
}

type mysqlVariable struct {
	Name  string `sql:"Variable_name"`
	Value string `sql:"Value"`
}

// 设置MySQL日志输出配置的函数
func setMySQLLogOutput() error {
	// 定义一个结构体变量用于存储查询到的变量名和值
	variable := mysqlVariable{}
	// 使用QueryRow查询数据库中名为'general_log_file'的变量
	row := db.QueryRow("SHOW VARIABLES LIKE 'general_log_file'")
	// 尝试扫描查询结果到variable结构体中
	if err := row.Scan(&variable.Name, &variable.Value); err != nil {
		// 如果扫描出错，返回错误
		return err
	}
	// 如果查询到的变量名是'general_log_file'，则更新全局变量logfile
	if variable.Name == "general_log_file" {
		logfile = variable.Value
	}

	// 执行SQL命令设置全局变量log_output为'FILE'，用于指定日志输出到文件
	if _, err := db.Exec("SET GLOBAL log_output = 'FILE'"); err != nil {
		// 如果设置出错，返回错误
		return err
	}
	// 执行SQL命令开启全局变量general_log，用于记录所有查询
	if _, err := db.Exec("SET GLOBAL general_log='ON'"); err != nil {
		// 如果开启出错，返回错误
		return err
	}
	// 如果所有操作成功，返回nil
	return nil
}

// cleanGenerakLog 函数用于关闭MySQL的通用日志功能，并清空日志文件
func cleanGenerakLog() error {
	// 尝试执行SQL命令关闭通用日志
	if _, err := db.Exec("SET GLOBAL general_log='OFF'"); err != nil {
		// 如果执行失败，返回错误
		return err
	}

	// 如果提供了日志文件路径
	if logfile != "" {
		// 使用os.Truncate将日志文件大小截断为0，即清空文件内容
		return os.Truncate(logfile, 0)
	}
	// 如果没有提供日志文件路径，返回nil表示成功
	return nil
}

type SQLInfo struct {
	count    int
	lastTime time.Time
	command  string
	sql      string
}

var (
	sqlMap   = make(map[string]*SQLInfo)
	mapMutex sync.Mutex
)

func isUnimportantSQL(sql string) bool {
	unimportantQueries := []string{"SET autocommit=0", "commit", "BEGIN", "ROLLBACK"}
	for _, q := range unimportantQueries {
		if strings.Contains(sql, q) {
			return true
		}
	}
	return false
}
func printResults() {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	var infos []*SQLInfo
	for _, info := range sqlMap {
		infos = append(infos, info)
	}

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].count > infos[j].count
	})

	fmt.Println("\nSQL执行统计:")
	fmt.Println("时间\t\t次数\t命令\tSQL")
	fmt.Println("----\t\t----\t----\t---")
	for _, info := range infos {
		fmt.Printf("%s\t%d\t%s\t%s\n",
			Yellow(info.lastTime.In(cstZone).Format("15:04:05")),
			info.count,
			Cyan(info.command),
			LightGreen(info.sql),
		)
	}
	fmt.Println() // 打印空行，增加可读性
}
func watchdog() {
	fmt.Println("start watchdog ...")

	if logfile == "" {
		log.Fatalln("general_log_file was empty :(")
	}
	fmt.Println("start watchdog ...", logfile)
	f, err := os.Open(logfile)
	if err != nil {
		log.Fatalf("Open '%s' error: %s", logfile, err)
	}
	defer f.Close()

	// 指向文件尾部
	_, err = f.Seek(0, 2)
	if err != nil {
		log.Fatalf("'%s' File.Seek(0,2) error: %s", logfile, err)
	}

	// make handle
	handle := func(line string) {
		if strings.Contains(line, "Execute") || strings.Contains(line, "Query") {
			lines := strings.Split(line, "\t")
			t, err := str2Time(lines[0], "2006-01-02T15:04:05Z")
			if err != nil {
				return
			}

			command := strings.Split(strings.TrimSpace(lines[1]), " ")[1]
			sql := strings.TrimSpace(lines[2])

			// 过滤不重要的SQL语句
			if isUnimportantSQL(sql) {
				return
			}

			key := command + sql
			mapMutex.Lock()
			if info, exists := sqlMap[key]; exists {
				info.count++
				info.lastTime = t
			} else {
				sqlMap[key] = &SQLInfo{
					count:    1,
					lastTime: t,
					command:  command,
					sql:      sql,
				}
			}
			mapMutex.Unlock()
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

LOOP:
	for {
		select {
		case <-quit:
			break LOOP
		case <-ticker.C:
			printResults()
		default:
			if err := linePrinter(f, handle); err != nil {
				log.Printf("linePrinter error: %s \n", err)
				break LOOP
			}
			time.Sleep(time.Millisecond * 550)
		}
	}

	// 在退出时打印最终结果
	printResults()
}

func linePrinter(r io.Reader, call func(string)) error {
	buf := make([]byte, 32*1024)
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		if c == 0 {
			return nil
		}
		for _, line := range bytes.Split(buf[:c], lineSep) {
			call(string(line))
		}
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
	}
}

// util part

func str2Time(timestr string, format string) (time.Time, error) {
	var (
		t   time.Time
		err error
	)
	t, err = time.Parse(format, timestr)
	if err != nil {
		return t, err
	}
	return t, nil
}

func isRoot() bool {
	return os.Geteuid() == 0
}
