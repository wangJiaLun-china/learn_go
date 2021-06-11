package main

import (
	"fmt"
	"net/rpc"
)

// RPC 使用了go语言特有的数据序列号gob，其他语言不能解析
func main() {
	// 1.rpc 连接服务器
	conn, err := rpc.Dial("tcp", ":8800")
	if err != nil {
		fmt.Println("rpc.Dial err: ", err)
	}
	defer conn.Close()
	fmt.Println("conn success...")
	// 2.调用远程函数
	var reply string
	err = conn.Call("hello.HelloWorld", "zhangsan", &reply)
	if err != nil {
		fmt.Println("conn.Call err: ", err)
	}
	fmt.Println("hello.HelloWorld.reply: ", reply)
}
