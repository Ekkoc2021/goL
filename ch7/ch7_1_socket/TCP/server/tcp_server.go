package main

import (
	"fmt"
	"net"
)

//创建端口,监听连接:net.Listen("tcp","localhost:9999")
//接收连接:listen.Accept()
//接收数据:n, err := con.Read(buf[:])
//读取数据:con.Write([]byte("你好,客户端!"))

func main() {
	//建立端口
	listen, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("建立端口出错:", err)
		return
	}

	for {
		//接收连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接收连接出错:", err)
			continue
		}
		//处理数据
		go process(conn)
	}
}
func process(con net.Conn) {
	defer con.Close()

	var buf [128]byte
	n, err := con.Read(buf[:]) //每次都会读满
	if err != nil {            //没有
		fmt.Println("读取错误:", err)
	}
	recvStr := string(buf[:n])
	fmt.Println("接收到的数据:", recvStr)
	con.Write([]byte("你好,客户端!"))

}
