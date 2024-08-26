package configMgmt

import (
	"fmt"
	"github.com/Adron/cobra-cli-samples/helper"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

// 删除配置键值对
func ConfigKeyValuePairDelete(key string) {
	// 使用Hack方式删除键
	DeleteKeyHack(key)
}

// 使用Hack方式删除键
func DeleteKeyHack(key string) {
	// 获取所有配置设置
	settings := viper.AllSettings()
	// 删除指定的键
	delete(settings, key)

	// 构建解析后的设置字符串
	var parsedSettings string
	for key, value := range settings {
		parsedSettings = fmt.Sprintf("%s\n%s: %s", parsedSettings, key, value)
	}

	// 将解析后的设置字符串转换为字节切片
	d1 := []byte(parsedSettings)
	// 将配置文件写入磁盘
	helper.HandleError(ioutil.WriteFile(viper.ConfigFileUsed(), d1, 0644))
}

// 更新配置键值对
func ConfigKeyValuePairUpdate(key string, value string) {
	// 写入键值对
	writeKeyValuePair(key, value)
}

// 添加配置键值对
func ConfigKeyValuePairAdd(key string, value string) {
	// 验证键值对
	if validateKeyValuePair(key, value) {
		// 验证未通过，记录日志
		log.Printf("Validation not met for %s.", key)
	} else {
		// 验证通过，写入键值对
		writeKeyValuePair(key, value)
	}
}

// 验证键值对
func validateKeyValuePair(key string, value string) bool {
	// 检查键和值是否为空
	if len(key) == 0 || len(value) == 0 {
		// 输出错误信息
		fmt.Println("The key and value must both contain contents to write to the configuration file.")
		// 返回true表示验证未通过
		return true
	}
	// 检查键是否已存在
	if findExistingKey(key) {
		// 输出错误信息
		fmt.Println("This key already exists. Create a key value pair with a different key, or if this is an update use the update command.")
		// 返回true表示验证未通过
		return true
	}
	// 返回false表示验证通过
	return false
}

// 写入键值对
func writeKeyValuePair(key string, value interface{}) {
	// 设置键值对
	viper.Set(key, value)
	// 写入配置文件
	err := viper.WriteConfig()
	// 处理错误
	helper.HandleError(err)
	// 输出成功信息
	fmt.Printf("Wrote the %s pair.\n", key)
}

// 查找存在的键
func findExistingKey(theKey string) bool {
	// 初始化是否存在键的标志
	existingKey := false
	// 遍历所有键
	for i := 0; i < len(viper.AllKeys()); i++ {
		// 如果找到匹配的键，设置标志为true
		if viper.AllKeys()[i] == theKey {
			existingKey = true
		}
	}
	// 返回是否存在键的标志
	return existingKey
}
