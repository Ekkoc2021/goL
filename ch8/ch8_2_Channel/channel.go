package main

import (
	"fmt"
	"time"
)

// 通道:可以在通道内放入数据,然后取出使用,可以用于两个goroutine之间通信,同时保证线程安全
// channel是一种数据类型
// channel 关闭后会读取到'0'值

// 管道基础操作:声明方式,赋值方式,输入值方式,获取值方式,关闭管道
func channeldemo1() {
	var ch chan int     //声明了一个int类型的通道,但是无法使用没有进行赋值
	ch = make(chan int) //为ch赋值
	go func() {
		fmt.Println("开始获取管道中的值....")
		i := <-ch //获取管道中的值,如果没有值则阻塞
		fmt.Println(i)
		fmt.Println("获取完毕!")
	}()

	time.Sleep(time.Second) //休眠1s
	ch <- 2                 //向管道输入值
	time.Sleep(time.Second) //休眠1s
	close(ch)               //关闭管道
}

// 缓冲管道:发送的数据可以在管道中存放,管道容量需要被声明
// 无缓冲管道:如上面的demo1使用的无缓冲管道,要求发送值,必须被接收不允许在管道停留否则报错
func channeldemo2() {
	ch := make(chan int, 2)
	ch <- 2
	ch <- 100
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 从有缓冲的管道中循环获取值:普通的for循环设置取值次数,使用管道中返回值进行标识,使用range
func channeldemo3() {
	ch := make(chan int, 2)
	ch <- 2
	ch <- 100
	go func() {
		for i := 0; i < 2; i++ { //管道中有两个值的时候会发送阻塞
			fmt.Println(<-ch)
		}

		for true {
			i, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Second)
	ch <- 99
	ch <- 939
	time.Sleep(time.Second)
	ch <- 1000
	ch <- 1001
	time.Sleep(time.Second)
	ch <- 1002
	ch <- 1003
	close(ch) //停止的发送数据
	go func() {
		for i := range ch { //没有值就不会进入循环
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Second)
}

// 单向通道:只可以发送,或者只允许接收
func channeldemo4() {
	ch := make(chan int, 2)
	var ch1 chan<- int //定义一个只可以发送的通道道
	ch1 = ch
	ch1 <- 10000

	var ch2 <-chan int
	ch2 = ch
	fmt.Println(<-ch2) //10000

	ch <- 1
	fmt.Println(<-ch2) // 1 管道赋值不是值类型赋值,这一点很重要
}
func main() {
	channeldemo1()
	channeldemo2()
	channeldemo3()
	channeldemo4()
}
