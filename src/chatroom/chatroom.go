package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("服务器启动成功，监听中...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	fmt.Println("建立连接成功!")

	go handler(conn)

}

// 处理具体业务
func handler(conn net.Conn) {
	fmt.Println("启动业务流程")
	// TODO 具体业务
	buf := make([]byte, 1024)

	// 读取客户端发送来的请求数据
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("服务器端接收数据:", string(buf[:cnt-1]))
}
