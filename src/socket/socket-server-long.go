package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 创建监听
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)
	// func Listen(network, address string) (Listener, error) {
	// net.Listen("tcp", ":8848") // ip不写默认本机
	linsten, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	// 需求：
	// 	server 可以接收多个连接，每个连接可以接收处理多轮数据请求
	//	每个连接可以接收处理多轮数据请求

	for {
		fmt.Println("监听中...")

		// Accept() (Conn, error)
		conn, err := linsten.Accept()
		if err != nil {
			fmt.Println("listener.Accept err :", err)
			return
		}

		fmt.Println("连接成功...")
		go handleFunc(conn)
	}
}

// 获取连接后的业务处理
func handleFunc(conn net.Conn) {
	// 创建一个容器，用于接收读取到的数据
	// 使用make创建字节切片
	buf := make([]byte, 1024)
	// Read(b []byte) (n int, err error)
	// cnt 真正读取的client数据长度
	cnt, err := conn.Read(buf)

	if err != nil {
		fmt.Println("conn.Read err: ", err)
	}

	fmt.Println("Client ======> Server, 长度：", cnt, ", 数据", string(buf[0:cnt]))

	// 数据处理
	// func ToUpper(s string) string
	upperData := strings.ToUpper(string(buf[0:cnt]))
	// Write(b []byte) (n int, err error)
	cnt, err = conn.Write([]byte(upperData))
	fmt.Println("Client <===== Server, 长度：", cnt, ", 数据: ", upperData)

	// 关闭连接
	_ = conn.Close()
}
