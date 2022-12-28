package main

import "net"

// 创建udp客户端: DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
// 读取数据: (c *conn) Read(b []byte) (int, error),这里UDPConn包装了一个conn
// 发送数据: (c *conn) Write(b []byte) (int, error)
// 关闭资源: close
func main() {
	listen, _ := net.DialUDP("udp", nil,
		&net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 8999},
	)
	defer listen.Close()
	listen.Write([]byte("你好服务器!"))
}
