package main

import "net/rpc"

type MyInterface interface {
	HelloWorld(string, *string) error
}

func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

func main() {

}
