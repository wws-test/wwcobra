package configMgmt

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {

	question := "介绍一下北京"
	roleDesc := "你是一位知识渊博的历史学家，擅长介绍中国各个城市的文化和历史。"

	answer, err := chatWithAI(question, roleDesc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Answer:", answer)
}
