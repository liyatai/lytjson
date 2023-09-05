package lytjson

import (
	"encoding/json"
	"fmt"
	"os"
)

func Write(url string, t any) {
	// 创建文件
	filePtr, err := os.Create(url)
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(t)
	if err != nil {
		fmt.Println("编码失败", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}
