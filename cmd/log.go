package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var taillogCmd = &cobra.Command{
	Use:   "taillog",
	Short: "Print the last 100 lines of a specified log file",
	Long: `The 'taillog' subcommand will print the last 100 lines of the specified log file. For example:

'<cmd> taillog /opt/metersphere/logs/api-test/ms-jmeter-run.log'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the path of the log file.")
			return
		}
		filePath := args[0]
		printLast100Lines(filePath)
	},
}

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

func init() {
	// 启用前缀匹配
	cobra.EnablePrefixMatching = true
	rootCmd.AddCommand(taillogCmd)
}
