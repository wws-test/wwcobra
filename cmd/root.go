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
	"github.com/Adron/cobra-cli-samples/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`,
}

// Execute函数用于执行rootCmd命令，如果执行出错则打印错误信息并退出程序
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init函数用于初始化配置
func init() {
	cobra.OnInitialize(initConfig) // 在初始化时调用initConfig函数
}

// initConfig函数用于初始化配置文件和环境变量
func initConfig() {
	configFile = ".cobra-cli-samples.yml" // 配置文件名
	viper.SetConfigType("yaml")           // 设置配置文件类型为yaml格式
	viper.SetConfigFile(configFile)       // 设置配置文件名

	viper.AutomaticEnv()                            // 读取环境变量
	viper.SetEnvPrefix("COBRACLISAMPLES")           // 设置环境变量前缀
	helper.HandleError(viper.BindEnv("API_KEY"))    // 绑定API_KEY环境变量
	helper.HandleError(viper.BindEnv("API_SECRET")) // 绑定API_SECRET环境变量
	helper.HandleError(viper.BindEnv("USERNAME"))   // 绑定USERNAME环境变量
	helper.HandleError(viper.BindEnv("PASSWORD"))   // 绑定PASSWORD环境变量

	if err := viper.ReadInConfig(); err == nil { // 读取配置文件，如果没有错误则打印使用的配置文件名
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
}
