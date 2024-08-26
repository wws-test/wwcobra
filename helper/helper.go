package helper

import "fmt"

// HandleError 函数用于处理错误信息
func HandleError(e error) {
	// 如果错误信息不为空，则打印错误信息
	if e != nil {
		fmt.Println(e)
	}
}
