package main

import (
	"fmt"
	"time"
)

// time这个库的一些简单api:实现定暂停
func main() {
	// 1.timer基本使用
	//timer1 := time.NewTimer(10 * time.Second)
	//fmt.Println("hello")
	//t1 := time.Now()
	//fmt.Printf("t1:%v\n", t1)
	//t2 := <-timer1.C//暂停10s
	//fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应1次
	//timer2 := time.NewTimer(time.Second)
	//for {
	//	<-timer2.C//只能使用一次,第二次异常
	//	fmt.Println("时间到")
	//}

	// 3.timer实现延时的功能
	//(1)
	//time.Sleep(time.Second)
	//(2)
	//timer3 := time.NewTimer(2 * time.Second)
	//<-timer3.C
	//fmt.Println("2秒到")
	//(3)
	//<-time.After(2*time.Second)
	//fmt.Println("2秒到")

	// 4.停止定时器
	//timer4 := time.NewTimer(2 * time.Second)
	//go func() {
	//	<-timer4.C
	//	fmt.Println("定时器执行了")
	//}()
	//b := timer4.Stop()
	//if b {
	//	fmt.Println("timer4已经关闭")
	//}
	//time.Sleep(3 * time.Second)

	// 5.重置暂停时间
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

}
