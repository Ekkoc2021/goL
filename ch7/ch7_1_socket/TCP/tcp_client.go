package main

import (
	"fmt"
	"net"
)

// 建立连接:net.Dial("tcp", "localhost:9999")
// 发送数据:conn.Write(b []byte) go语言字符串转byte切片直接强转即可[]byte(string)
// 接收数据:n, err := conn.Read(buf[:])    Read(b []byte) (n int, err error),使用一个切片接收
// 关闭连接:conn.Close()
func main() {

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer conn.Close() //关闭服务,无论如何都要关闭
	for {
		//发送数据:conn.Write(b []byte) go语言字符串转byte数组直接强转即可[]byte(string)
		_, err = conn.Write([]byte("你好,服务器!"))
		if err != nil {
			return
		}

		//接收数据
		buf := [512]byte{}
		n, err := conn.Read(buf[:]) //读取到的长度
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))

	}

}
