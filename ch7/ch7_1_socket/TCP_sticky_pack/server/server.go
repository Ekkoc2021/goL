package main

import (
	"fmt"
	"io"
	"net"
)

// 我们发现:接收到的客户端数据,表示分条接收,有一些黏在了一起
// 为什么?
// 主要是因为nagle算法导致的(提高网络传输效率)
// 如果接收端接收不及时,会将数据放在缓冲区和上一次的接在一起
// 同时发送端也会判断是否还有数据发送,将多次发送数据接在一起发送
func process(conn net.Conn) {
	defer conn.Close()
	//reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
