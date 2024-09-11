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
	"github.com/spf13/cobra"
)

// configCmd represents the config command
// fileName:config.go

// ConfigCmd 是一个用于管理配置的子命令，可以与其他子命令 'add', 'update', 'view', 和 'delete' 结合使用。
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "配置-config子命令用于配置管理.",
	Long: `“'config子命令用于配置管理。它可以与其他子命令add'、'update'、'view和'delete'结合使用。例如：

Cobra是一个用于Go的CLI库，赋予应用程序力量。
这个应用程序是一个用于快速生成所需文件的工具
以快速创建一个Cobra应用程序。”`,
}

func init() {
	rootCmd.AddCommand(ConfigCmd)
	// 为 configCmd 添加持久标志，用于添加到配置的键值对的键和值。
	ConfigCmd.PersistentFlags().StringP("key", "k", "", "The key for the key value set to add to the configuration.")
	ConfigCmd.PersistentFlags().StringP("value", "v", "", "The value for the key value set to add to the configuration.")
}
