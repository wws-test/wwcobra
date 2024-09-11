package cmd

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func printLast100Lines(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 100)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) > 100 {
			lines = lines[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

var taillogCmd = &cobra.Command{
	Use:   "taillog",
	Short: "打印指定日志文件的最后 100 行",
	Long: `‘taillog’ 子命令将打印指定日志文件的最后 100 行。例如：

'<cmd> taillog /opt/metersphere/logs/api-test/ms-jmeter-run.log'`,
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string

		if len(args) < 1 {
			prompt := promptui.Prompt{
				Label:   "请输入日志文件的路径",
				Default: "/opt/metersphere/logs/api-test/ms-jmeter-run.log",
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
			filePath = result
		} else {
			filePath = args[0]
		}

		printLast100Lines(filePath)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 这里可以根据需要返回自动补全的候选列表
		// 例如，返回当前目录下的所有文件和目录
		files, _ := filepath.Glob(toComplete + "*")
		return files, cobra.ShellCompDirectiveDefault
	},
}

//var taillogCmd = &cobra.Command{
//	Use:   "taillog",
//	Short: "Print the last 100 lines of a specified log file",
//	Long: `The 'taillog' subcommand will print the last 100 lines of the specified log file. For example:
//
//'<cmd> taillog /opt/metersphere/logs/api-test/ms-jmeter-run.log'`,
//	Run: func(cmd *cobra.Command, args []string) {
//		if len(args) < 1 {
//			prompt := promptui.Prompt{
//				Label: "Please provide the path of the log file",
//			}
//
//			result, err := prompt.Run()
//			if err != nil {
//				fmt.Printf("Prompt failed %v\n", err)
//				return
//			}
//
//			filePath := result
//			printLast100Lines(filePath)
//		} else {
//			filePath := args[0]
//			printLast100Lines(filePath)
//		}
//	},
//	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
//		// 这里可以根据需要返回自动补全的候选列表
//		// 例如，返回当前目录下的所有文件和目录
//		files, _ := filepath.Glob(toComplete + "*")
//		return files, cobra.ShellCompDirectiveDefault
//	},
//}

func init() {
	// 启用前缀匹配
	cobra.EnablePrefixMatching = true
	rootCmd.AddCommand(taillogCmd)
}
