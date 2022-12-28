package main

import (
	"fmt"
	"net/http"
)

// go能轻松完成发送http请求
func main() {
	resp, _ := http.Get("http://localhost:8080/go")
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	//接收服务端数据
	data := make([]byte, 1024)
	resp.Body.Read(data)
	fmt.Println(string(data))
}
