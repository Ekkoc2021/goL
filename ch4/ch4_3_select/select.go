package main

import (
	"fmt"
	"time"
)

func main() {
	selectStatement()
}

// select语句
func selectStatement() {
	fmt.Println("main start...")
	//select用于监听通道是否写入数据
	var c1 = make(chan int, 1) //创建接收int的chanel
	var c2 = make(chan int, 2)
	go func() { //开启一个协程,可以暂时理解为线程,但是和线程有区别,协程属于线程
		fmt.Println("func start..")
		loop := true
		for loop { //循环监听,我们重复运行程序发现打印顺序是不一致,从而该监听是随机的
			select {
			case msg, ok := <-c1: //定义两个变量,当通道写入数据时,用msg读取数据,ok表示读取成功与否
				fmt.Println("通道1")
				if ok {
					fmt.Println(msg)
				}
			case msg, ok := <-c2:
				if msg == 90 { //为了能看到func end 添加一个跳出循环的条件
					loop = false
					break
				}
				fmt.Println("通道2")
				if ok {
					fmt.Println(msg)
				}
			default: //一般default是必须的,否则当没有任何数据写入一直阻塞会耗费cpu
				time.Sleep(time.Second) //太快了,会打印太多
				fmt.Println("没有监听到!")
			}
		}
		fmt.Println("func end")
		//如果没有任何一个通道接收到数据,那么会阻塞知道有通道接收数据
	}()
	c2 <- 2 //向通道2写入数据2
	c1 <- 1
	time.Sleep(time.Second) //main停一下等待协程监听数据,不然太快了
	c2 <- 90
	time.Sleep(time.Second) //再停一下让协程再监听一下
	fmt.Println("main end")
}
