package main

import (
	"fmt"
)

//函数
/*
函数不可以定义在函数体内,函数体内可以使用匿名函数
go中的函数不支持重载

*/

// -------------函数类型
// go中的函数也是一种类型
// 函数类型,当一个函数的参数类型顺序相同且返回值类型个数相同,这个函数属于统一类型
// 当函数作为类型时,参数名称可以省略
type getSum func([]int) int //为这样的函数类型明一个别的名称为getSum

func sum(a []int) int { //这个函数和getSum函数类型一致
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func funcTypeDemo(s getSum, prefix string, arr []int) { //传入一个getSum函数然后使用
	fmt.Printf("%s%d", prefix, s(arr))

	type student struct {
		study func() //为学生这个结构体定义了一个函数类型的变量,相当于一个接口,具体的函数我们没有定义
	}
}

// -------------------参数
// 当定义函数的形式参数时,如果多个形式参数类型相同且连续时,可以只声明最后一个类型
// 一般实参通过值的方式传递,除非实参是应用类型或者指针,slice,map,function,channel等
// 可以传递多个参数,但是一般放在参数的最后:args...T 这样定义意味着所有T类型参数都会统一放在args中,args的类型是T的切片
// slice可以展开传递给多个参数的形式参数,s...
func argsDemo(a int, m ...string) {
	fmt.Println(a)
	fmt.Printf("类型是:%T", m)
	m[0] = "我修改了第一个值!"
	for v := range m {
		fmt.Println(v)
	}
}

// -------------------返回值
// 1,go中函数可以有多个返回值:func name(...)(return1 T,return2 T,...){}
// 2,当有多个返回值时,用一个变量接收那么只会接收到第一个返回值
// 3,多个返回值,但有的不需要使用该返回值,但不得不接收时可以用'_'表示忽略
// 4,go中即时return语句没有返回所有的值也不会报错,这种称为裸返回
// 5,多个返回值可以直接作为其他函数的参数使用,只要类型对得上
// 6,函数有返回值是return是必须的
// 7,go中返回值可以被命名,这样是有意义的可以作为文档使用,命名时或者多个返回值时必须使用小括号
// 8,命名返回参数可以被看做与形参相似的局部变量,即使是裸返回,只要函数体内定义了与返回参数相同的名称也会被隐式返回
func returnDemo1() (state int, msg string) {
	msg = "这样你也能收到?" //返回参数已经定义好了,可以直接使用然后被返回
	// 低版本的go好像可以屏蔽掉返回参数中的定义的变量
	return 1, ""
}

// 9,命名返回参数允许defer延迟调用通过闭包读取和修改
// 执行顺序:最后只差return时,执行defer中代码
func returnDemo2() (msg string) {
	defer func() {
		msg += " defer到此处一游!"
		fmt.Println(msg)
	}()
	msg = "返回吧!"
	return
}

// ----------------匿名函数
func anonymityFuncDemo() {
	anoF := func() {
		fmt.Println("匿名函数")
	}
	anoF()                                   //调用
	func() { fmt.Println("hello world!") }() //直接调用
	//还可以这样用
	anoF2 := []func(x int){
		func(x int) {
			fmt.Println(x + 1)
		},
		func(x int) {
			fmt.Println(x + 2)
		},
	}
	anoF2[0](2)
	anoF2[1](1000)
}
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T \n", a)
	x := a[:]
	funcTypeDemo(sum, "1-5的和是", x)

	str := []string{"a", "b", "c", "d"}
	argsDemo(2, str...) //slice可以像这样传递给args,这样告诉args形参我们也是一个切片你可以直接使用
	fmt.Println(str[0])

	_, msg := returnDemo1()
	fmt.Println(msg)

	anonymityFuncDemo()
}
