package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 开启goroutine只需要一个go关键词:go 执行函数
// 主程序接收goroutine也会结束
func goroutineDemo1() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 2)
}

// 我们可以同下面方式,使得主线程等待所有goroutine运行结束
func goroutineDemo2() {
	var wg sync.WaitGroup
	go func() {
		fmt.Println("快递结束吧!")
		runtime.Goexit()
		fmt.Println("qnm")
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go func() {
			defer wg.Done() // goroutine结束就登记-1
			fmt.Println("Hello Goroutine!", i)
		}()
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

//goroutine可以实现线程一样的功能,但是实际上goroutine并不是线程,他和os的线程有着很大区别
//os线程有固定的内存大小2mb,其他语言开启的线程就是os线程
//2mb的固定占用,1g内存只够开启512个,1w个线程差不多20g的内存
//goroutine的内存大小初始化为4k可以动态增加,4k意味着2mb就可以开启512个goroutine
//一个os线程可以执行多个goroutine

//执行goroutine时,可以使用runtime内置的一些方法,进行cpu的再分配
//runtime.Gosched():让出当前cpu时间片,轮到该goroutine执行时,先让其他的先执行
//runtime.Goexit():主动结束当前goroutine
//runtime.GOMAXPROCS(i int):确定当前代码使用的cpu核心数量,默认值是机器的核心数量

func main() {
	runtime.GOMAXPROCS(2)
	goroutineDemo1()
	goroutineDemo2()
}
