package main

import (
	"fmt"
	"time"
)

// defer:用于注册延迟调用,当程序发生异常就提前进入defer语句,直到异常在defer中被处理,程序才会恢复执行
//这些调用直到return前才被执行,可以用于清理资源,defer的调用时机很重要,一定要深刻理解

// defer语句是在函数中所有语句执行完毕只剩下返回的时候执行
func deferdemo() (i int) {
	i = 0
	defer func() {
		fmt.Println(i) //执行的时候记录的环境内i的地址
	}()
	return 100 //这里retrun时还有一句 i=100 ,所以最后输出100
}

// 多个defer执行顺序,是从下往上,如果前面将资源释放,后面语句可能无法执行
func deferdemo1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
func deferdemo2() {
	for i := 0; i < 5; i++ {
		defer func() { fmt.Println(i) }()
		//函数中会记录i的值,后续执行的时候i值已经变更为5,类似于闭包
	}
}

// 外界变量作为参数传递时,会预先将值保存起来,但是函数没有执行
func deferdemo3() {
	for i := 0; i < 5; i++ {
		defer func(i int) { fmt.Println(i) }(i) //4 3 2 1 0 和第一个类似
	}
}

// defer调用顺序先进后出,由下往上执行,而且即使中间发生错误也不会停止
func deferdemo4(x int) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer func() { //一开始并没有执行该函数,从而1/x没有处理
		fmt.Println(1 / x)
	}()
	defer fmt.Println("4")
}
func deferdemo5(x int) { //4,抛出异常
	defer fmt.Println("1")   //3,准备捕获异常,没有捕获
	defer fmt.Println("2")   //2,准备捕获异常,没有捕获
	defer fmt.Println(1 / x) //1,没有执行,赋值过程中go会将语句进行运算,从而抛出异常
	defer fmt.Println("4")
}

// defer对性能有影响
func deferdemo6() {
	start1 := time.Now()
	sum := 0
	for i := 0; i < 1000000; i++ {
		func(i2 *int) { *i2 = *i2 + i }(&sum)
	}
	fmt.Println(time.Since(start1)) //2.0939ms

	start1 = time.Now()
	for i := 0; i < 1000000; i++ {
		defer func(i2 int) { sum = sum + i2 }(i)
	}
	fmt.Println(time.Since(start1)) //219.1542ms
	//差距在100倍左右,而下面一个也仅仅好像只是比上面多了一个循环
}
func main() {
	deferdemo()
	//deferdemo1() //43210
	//deferdemo2()
	//deferdemo3()
	//deferdemo4(0) //4 2 1 有语言经验反而不好理解,函数内抛出异常,马上终止该程序
	//deferdemo5(0) //2 1
	deferdemo6()

}
