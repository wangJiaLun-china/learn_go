package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		// writer ==》 通过writer将数据返回客户端
		// request ==》 客户端请求数据
		fmt.Println(request)
		io.WriteString(writer, "user 信息")
	})

	// func ListenAndServe(addr string, handler Handler) error {
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Print("http start failed, err: ", err)
		return
	}
}
