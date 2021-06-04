package main

import (
	"encoding/json"
	"fmt"
)

/**
小写字母开头字段会忽略，不参与json编解码
后面配置 `json:"-"` 忽略
配置 `json:"Age_int"` 变成json中别名
配置 `json:"Age_int,string"` 变成json中别名,并转换类型
配置 `json:"Age_int,omitempty"` 变成json中别名,如果为空就不编解码
*/
type User struct {
	Id   int `json:"-"`
	Age  int `json:"Age_int,string"`
	Name string
}

func main() {
	zhangsan := User{
		Id:   1,
		Name: "zhangsan",
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
