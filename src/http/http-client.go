package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}

	// func (c *Client) Get(url string) (resp *Response, err error) {
	resp, err := client.Get("http://localhost:8080/user")
	if err != nil {
		fmt.Println("client.Get error:", err)
		return
	}

	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	fmt.Println(ct, date)
	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read responseBody err: ", err)
		return
	}
	fmt.Println("body string:", string(body))

}
