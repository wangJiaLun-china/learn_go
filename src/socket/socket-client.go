package main

import (
	"fmt"
	"net"
)

func main() {
	var conn, err = net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	fmt.Println("client 与 server 连接成功！")
	sendData := []byte("helloWorld")

	cnt, err := conn.Write(sendData)
	if err != nil {
		fmt.Println("conn.write err:", err)
		return
	}

	fmt.Println("Client ===> server cnt:", cnt, ",data:", string(sendData))

	// 接收服务器返回数据
	buf := make([]byte, 1024)

	cnt, err = conn.Write(sendData)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("client <==== Server, cnt:", cnt, ",data:", string(buf[cnt]))

	conn.Close()
}
