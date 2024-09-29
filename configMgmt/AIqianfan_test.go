package configMgmt

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {

	question := "介绍一下北京"
	roleDesc := "以json模式输出。"

	answer, err := chatWithAI(question, roleDesc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Answer:", answer)
}
