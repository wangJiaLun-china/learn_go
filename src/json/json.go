package main

import (
	"encoding/json"
	"fmt"
)

/**
小写字母开头字段会忽略，不参与json编解码
*/
type User struct {
	Id   int
	Age  int
	Name string
}

func main() {
	zhangsan := User{
		1,
		20,
		"zhangsan",
	}
	// func Marshal(v interface{}) ([]byte, error) {
	encodeInfo, err := json.Marshal(&zhangsan)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	fmt.Println(string(encodeInfo))

	// 反序列号
	var zhangsan2 User
	if err := json.Unmarshal(encodeInfo, &zhangsan2); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Println(zhangsan2.Age)
	fmt.Println(zhangsan2.Name)
}
