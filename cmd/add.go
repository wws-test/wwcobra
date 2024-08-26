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

// addCmd represents the add command
// addCmd 是一个用于添加键值对到应用程序配置文件的子命令
var addCmd = &cobra.Command{
	Use:   "add",                               // 使用说明
	Short: "‘add’ 子命令会将传入的键值对添加到应用程序配置文件中。\".", // 简短说明
	Long: `" 'add' 子命令向应用程序配置文件添加键值对。例如：

'<cmd> config add --key theKey --value "可以是一系列事物的值。""'.`, // 详细说明
	Run: func(cmd *cobra.Command, args []string) { // 当执行该命令时调用的函数
		key, _ := cmd.Flags().GetString("key")       // 获取命令行参数中的key值
		value, _ := cmd.Flags().GetString("value")   // 获取命令行参数中的value值
		configMgmt.ConfigKeyValuePairAdd(key, value) // 调用configMgmt包中的ConfigKeyValuePairAdd函数，将key value添加到配置文件中
	},
}

func init() {
	configCmd.AddCommand(addCmd) // 将addCmd作为configCmd的子命令
}
