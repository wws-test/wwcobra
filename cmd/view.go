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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viewCmd represents the view command
// 定义名为viewCmd的命令变量，用于展示键的列表和值的映射
var viewCmd = &cobra.Command{
	Use:   "view",                                                                       // 命令使用的名称
	Short: "The 'view' subcommand will provide a list of keys and a map of the values.", // 命令的简短描述
	Long: `The 'view' subcommand will provide a list of keys and a map of the values. For example:

'<cmd> config view'`, // 命令的详细描述
	Run: func(cmd *cobra.Command, args []string) { // 命令执行时调用的函数
		fmt.Printf("** All keys including environment variables for CLI.\n") // 打印所有包括CLI环境变量的键
		fmt.Printf("%s\n\n", viper.AllKeys())                                // 打印所有键并换行

		settings := viper.AllSettings()                        // 获取所有设置
		fmt.Printf("** Configuration file keys and values.\n") // 打印配置文件的键和值
		for i, v := range settings {                           // 遍历所有设置
			fmt.Printf("%v: %v\n", i, v) // 打印键和对应的值
		}
	},
}

func init() {
	configCmd.AddCommand(viewCmd) // 将viewCmd命令添加到configCmd中
}
