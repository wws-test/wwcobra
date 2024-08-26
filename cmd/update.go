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
	"github.com/Adron/cobra-cli-samples/configMgmt"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
// fileName:update.go

// updateCmd 是一个用于更新应用程序配置文件中现有数据的键值对的子命令
var updateCmd = &cobra.Command{
	Use:   "update",                                                                                                                            // 使用update子命令
	Short: "The 'update' subcommand will update a passed in key value pair for an existing set of data to the application configuration file.", // "update"子命令将更新应用程序配置文件中现有数据的传入键值对
	Long: `The 'update' subcommand updates a key value pair, if the key value pair already exists it is updated, if it does
not exist then the passed in values are added to the application configuration file. For example:

'<cmd> config add --key theKey --value "the new value which will be updated for this particular key value pair."'.`, // "update"子命令更新键值对，如果键值对已存在，则更新它，如果不存在，则将传入的值添加到应用程序配置文件中。例如：'<cmd> config add --key theKey --value "the new value which will be updated for this particular key value pair."'
	Run: func(cmd *cobra.Command, args []string) { // 运行update子命令时执行的函数
		key, _ := cmd.Flags().GetString("key")          // 获取命令行中的key标志值
		value, _ := cmd.Flags().GetString("value")      // 获取命令行中的value标志值
		configMgmt.ConfigKeyValuePairUpdate(key, value) // 调用configMgmt包中的ConfigKeyValuePairUpdate函数，更新键值对
	},
}

func init() {
	configCmd.AddCommand(updateCmd) // 将updateCmd子命令添加到configCmd命令中
}
