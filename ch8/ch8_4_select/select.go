package main

import (
	"fmt"
	"time"
)

// 多路复用:类似于switch语句,同时判断通道是否有数据输入
func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}
func demo1() {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go test1(output1)
	go test2(output2)
	// 用select监控
	select { //只要有一个通道有数据输入,就接收,同时结束语句
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}

// 我们知道当通道中数据满时再写入数据就会出现错误,使用select语句在通道满时,不会写入数据
func write(ch chan string) {
	for i := 0; i < 3; i++ {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
	}
	close(ch)
}
func demo2() {
	// 创建管道
	output1 := make(chan string, 1)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		time.Sleep(time.Second) //休眠等待填入数据
		fmt.Println("res:", s)
	}
}
func main() {
	demo1()
	demo2()
}
