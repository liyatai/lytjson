package lytjson

import (
	"encoding/json"
	"fmt"
	"os"
)

// Type返回值类型
func Read(url string, t any) {
	filePtr, err := os.Open(url)
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]", err.Error())
	}
	defer filePtr.Close()
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(t)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")

	}

}
