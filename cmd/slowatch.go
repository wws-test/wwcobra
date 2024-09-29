package cmd

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed kyanos
var kyanosBinary embed.FS

//func extractAndRunKyanos(option string, port string) error {
//	tmpFile := filepath.Join(os.TempDir(), "kyanos")
//	data, err := kyanosBinary.ReadFile("kyanos")
//	if err != nil {
//		return fmt.Errorf("failed to read embedded kyanos binary: %v", err)
//	}
//
//	err = ioutil.WriteFile(tmpFile, data, 0755)
//	if err != nil {
//		return fmt.Errorf("failed to write kyanos binary to tmp: %v", err)
//	}
//
//	var cmd *exec.Cmd
//	if option == "mysql" {
//		// 执行 ./kyanos stat mysql --side client -i 5 -m n -l 10 -g conn --local-ports=<port>
//		cmd = exec.Command(tmpFile, "stat", "mysql", "--side", "client", "-i", "5", "-m", "n", "-l", "10", "-g", "conn", "--local-ports="+port)
//	} else if option == "http" {
//		// 执行 ./kyanos stat http --side client -i 5 -m n -l 10 -g conn --local-ports=<port>
//		cmd = exec.Command(tmpFile, "stat", "http", "--side", "client", "-i", "5", "-m", "n", "-l", "10", "-g", "conn", "--local-ports="+port)
//	}
//
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		return fmt.Errorf("failed to execute kyanos: %v", err)
//	}
//
//	fmt.Printf("kyanos output:\n%s\n", string(output))
//	os.Remove(tmpFile)
//
//	return nil
//}

// // 提取并执行 kyanos 的函数
//
//	func extractAndRunKyanos(args []string) error {
//		// 提取嵌入的二进制文件
//		tmpFile := filepath.Join(os.TempDir(), "kyanos")
//		data, err := kyanosBinary.ReadFile("kyanos")
//		if err != nil {
//			return fmt.Errorf("无法读取嵌入的 kyanos 二进制文件: %v", err)
//		}
//
//		// 将二进制文件写入临时目录
//		err = ioutil.WriteFile(tmpFile, data, 0755)
//		if err != nil {
//			return fmt.Errorf("无法将 kyanos 二进制文件写入临时文件: %v", err)
//		}
//
//		// 构建要执行的命令
//		cmd := exec.Command(tmpFile, args...)
//
//		// 获取命令的标准输出和错误输出的管道
//		stdoutPipe, err := cmd.StdoutPipe()
//		if err != nil {
//			return fmt.Errorf("无法获取标准输出管道: %v", err)
//		}
//		stderrPipe, err := cmd.StderrPipe()
//		if err != nil {
//			return fmt.Errorf("无法获取错误输出管道: %v", err)
//		}
//
//		// 启动命令
//		if err := cmd.Start(); err != nil {
//			return fmt.Errorf("启动 kyanos 失败: %v", err)
//		}
//
//		// 创建一个 Goroutine 读取标准输出并打印到控制台
//		go func() {
//			scanner := bufio.NewScanner(stdoutPipe)
//			for scanner.Scan() {
//				fmt.Printf("[STDOUT] %s\n", scanner.Text()) // 实时打印输出
//			}
//		}()
//
//		// 创建一个 Goroutine 读取标准错误并打印到控制台
//		go func() {
//			scanner := bufio.NewScanner(stderrPipe)
//			for scanner.Scan() {
//				fmt.Printf("[STDERR] %s\n", scanner.Text()) // 实时打印错误输出
//			}
//		}()
//
//		// 等待命令完成
//		if err := cmd.Wait(); err != nil {
//			return fmt.Errorf("执行 kyanos 失败: %v", err)
//		}
//
//		// 删除临时文件
//		os.Remove(tmpFile)
//
//		return nil
//	}
func extractAndRunKyanos(args []string) error {
	// 提取嵌入的二进制文件
	tmpFile := filepath.Join(os.TempDir(), "kyanos")
	data, err := kyanosBinary.ReadFile("kyanos")
	if err != nil {
		return fmt.Errorf("无法读取嵌入的 kyanos 二进制文件: %v", err)
	}

	// 将二进制文件写入临时目录
	err = ioutil.WriteFile(tmpFile, data, 0755)
	if err != nil {
		return fmt.Errorf("无法将 kyanos 二进制文件写入临时文件: %v", err)
	}

	// 构建要执行的命令
	cmd := exec.Command(tmpFile, args...)

	// 将标准输出和标准错误直接指向控制台
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 启动命令
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("启动 kyanos 失败: %v", err)
	}

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("执行 kyanos 失败: %v", err)
	}

	// 删除临时文件
	os.Remove(tmpFile)

	return nil
}

var mode string
var port string

//	var slowCmd = &cobra.Command{
//		Use:   "slow",
//		Short: "查询慢请求sql",
//		Long:  `‘slowdebug’ 子命令将查询慢请求sql，每5秒输出请求响应在网络中的耗时最长的前10个HTTP连接`,
//		Run: func(cmd *cobra.Command, args []string) {
//			var option string
//			var userPort string
//
//			// 如果用户未提供选项，则使用 Select 提示用户选择 "mysql" 或 "http"
//			if len(args) < 1 {
//				prompt := promptui.Select{
//					Label: "选择 http 或者 mysql",
//					Items: []string{"http", "mysql"},
//				}
//
//				_, result, err := prompt.Run()
//				if err != nil {
//					fmt.Printf("选择失败: %v\n", err)
//					return
//				}
//				option = result
//			} else {
//				option = args[0]
//			}
//
//			// 选择后进一步让用户输入端口号
//			portPrompt := promptui.Prompt{
//				Label: "请输入端口号",
//				Validate: func(input string) error {
//					if strings.TrimSpace(input) == "" {
//						return fmt.Errorf("端口号不能为空")
//					}
//					// 检查输入是否为数字
//					if _, err := strconv.Atoi(input); err != nil {
//						return fmt.Errorf("端口号必须是数字")
//					}
//					return nil
//				},
//			}
//
//			// 运行端口号输入提示
//			userPort, err := portPrompt.Run()
//			if err != nil {
//				fmt.Printf("输入端口号失败: %v\n", err)
//				return
//			}
//
//			fmt.Printf("您选择了: %s\n输入的端口号是: %s\n", option, userPort)
//
//			// 根据用户选择执行不同的命令，端口号作为参数传递
//			var execErr error
//			switch strings.ToLower(option) {
//			case "mysql":
//				// 执行 ./kyanos stat mysql --side client -i 5 -m n -l 10 -g conn --local-ports=<userPort>
//				execErr = extractAndRunKyanos("mysql", userPort)
//			case "http":
//				// 执行 ./kyanos stat http --side client -i 5 -m n -l 10 -g conn --local-ports=<userPort>
//				execErr = extractAndRunKyanos("http", userPort)
//			default:
//				fmt.Println("无效的选项，请选择 'mysql' 或 'http'")
//				return
//			}
//
//			if execErr != nil {
//				fmt.Printf("执行 kyanos 失败: %v\n", execErr)
//			}
//		},
//
//		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
//			files, _ := filepath.Glob(toComplete + "*")
//			return files, cobra.ShellCompDirectiveDefault
//		},
//	}
var slowCmd = &cobra.Command{
	Use:   "slowdebug",
	Short: "查询慢请求sql",
	Run: func(cmd *cobra.Command, args []string) {
		// 构建参数列表s
		args = []string{"stat", mode, "--side", "client", "-i", "5", "-m", "p", "-s", "10", "-g"}
		if port != "" {
			args = append(args, "--local-ports="+port)
		}

		// 提取并执行嵌入的 kyanos
		err := extractAndRunKyanos(args)
		if err != nil {
			fmt.Printf("执行 kyanos 失败: %v\n", err)
		}
	},
}

func init() {
	slowCmd.Flags().StringVarP(&mode, "mode", "m", "", "运行模式 (mysql 或 http)")
	slowCmd.Flags().StringVarP(&port, "port", "p", "", "自定义端口号")
	// 启用前缀匹配
	cobra.EnablePrefixMatching = true
	rootCmd.AddCommand(slowCmd)

}
