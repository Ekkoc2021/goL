package main

import (
	"errors"
	"fmt"
)

// 异常:了解异常处理机制,同时掌握自定义异常
// 和其他语言比起来go中处理异常的方式有些怪异,大部分语言处理异常的结构是相似的
// panic(msg):主动抛出异常,msg是任意类型
// recover() msg:获取抛出的异常并将其返回,msg是任意类型
// defer语句,发生异常立即进入该语句,并且倒序执行异常前面的所有的defer语句

// 当使用recover捕获到异常时,程序会接着执行
func errdemo1() {
	defer func() {
		fmt.Println("我想看看我在什么时候执行!")
	}()
	defer func() {
		e := recover()
		if e != nil {
			fmt.Printf("%T\n", e)
		}
		fmt.Println("捕获结束")
	}() //发生异常后,我们在这里捕获
	defer recover() //异常有被捕获并且被使用才会终止,无效捕获
	fmt.Println("hello ")
	panic(1)
	fmt.Println("world!")
}

// 定义错误信息:errors.New("msg")返回一个error
func errdemo2() {
	defer func() {
		e := recover()
		fmt.Println(e)
	}()
	err := errors.New("msg")
	panic(err)
	fmt.Println("hjello")
}

// 当代码
func errdemo3(x int) {
	func() {
		defer func() {
			e := recover()
			fmt.Println(e)
		}()
		a := 9 / x
		fmt.Println(a)
	}()
	fmt.Println("hello world!")
}
func main() {
	errdemo1()
	errdemo2()
	var a int
	fmt.Scan(&a)
	errdemo3(a) //不能在编译阶段就使得传入的a=0,否则编译不通过
}
