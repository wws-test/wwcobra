/*
Copyright © 2019 Adron Hall <adron@thrashingcode.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/Adron/cobra-cli-samples/configMgmt"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
// deleteCmd 是一个 cobra.Command 类型的变量，用于表示 'delete' 子命令，用于从配置文件中删除键值对。
var deleteCmd = &cobra.Command{
	// Use 用于指定命令的使用方式。
	Use: "delete",
	// Short 用于提供命令的简短描述。
	Short: "The 'delete' subcommand removes a key value pair from the configuration file.",
	// Long 用于提供命令的详细描述。
	Long: `The 'delete' subcommand removes a key value pair from the configuration file.`,
	// Run 为命令执行时的操作，其中包含了获取参数并调用相应函数的逻辑。
	Run: func(cmd *cobra.Command, args []string) {
		// 从命令行参数中获取键值对的键。
		key, _ := cmd.Flags().GetString("key")
		// 打印删除操作的提示信息。
		fmt.Printf("\n\n    **** Deleting key: %s ****\n\n", key)
		// 调用配置管理包中的函数，删除指定的键值对。
		configMgmt.ConfigKeyValuePairDelete(key)
	},
}

func init() {
	configCmd.AddCommand(deleteCmd)
}
