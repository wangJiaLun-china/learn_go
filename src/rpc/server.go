package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义类对象
type World struct {
}

// 绑定类方法
func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + " hello!"
	return nil
}

func main() {
	// 1.注册RPC服务， 绑定对象方法
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("rpc.RegisterName err: ", err)
	}
	// 2.设置监听
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
	}
	defer listener.Close()

	fmt.Println("Listening...")
	// 3.建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err: ", err)
	}
	defer conn.Close()
	fmt.Println("conn success")
	// 4.绑定服务
	rpc.ServeConn(conn)
}
