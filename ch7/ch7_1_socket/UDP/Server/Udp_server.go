package main

import (
	"fmt"
	"net"
)

// 创建监听的udp: ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
// 读取数据: (c *conn) Read(b []byte) (int, error),这里UDPConn包装了一个conn
// 发送数据: (c *conn) Write(b []byte) (int, error)
// 关闭资源: close
// 其实就一点,发送数据要用byte切片装,使用字符串转byte切片的api
// 如果对参数不够了解,直接点进去看看原源码的结构
func main() {
	listen, _ := net.ListenUDP("udp",
		&net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 8999},
	)
	defer listen.Close()
	for {
		var data [1024]byte
		listen.Read(data[:])
		s := string(data[:])
		fmt.Println(s)
	}

}
