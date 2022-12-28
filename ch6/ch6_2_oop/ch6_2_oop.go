package main

import (
	"fmt"
)

//go语言不支持类,但是有和类功能类似语法:如动态绑定,和继承
//前面提到的匿名字段可以很好的实现继承
//面向接口编程是一个很好的思想,go语言尽管没有类的概念但是支持接口,通过接口能够很好实现类似动态绑定
//接口用于定义方法集合而不作具体实现,在go语言中,interface也是一种类型
//go剔除了类这个概念,但是保留面向对象的一些特性

//我接触面向对象的语言时间最多,用面向过程思想进行编程时,总是想往面向对象上靠
//当我在go语言的目录没有看到类这一章节的内容时,多少会有些排斥,很高兴发现go可以实现继承,动态绑定

// 接口定义:记住接口一种类型
// 实现接口不需要其他的语法,你只要将其方法全部实现即可
// 接口类型可以直接接收指针类型的实现类,go语言做了优化
type Animal interface {
	//定义了两个方法
	Say()
	Eat()
}

type cat struct {
}

func (cat) Say() {
	fmt.Println("喵")
}
func (cat) Eat() {
	fmt.Println("鱼")
}
func (cat) behavior() {
	fmt.Println("爬树")
}
func InterfaceDemo1() {
	var a Animal = cat{} //如果没有实现所有的接口,是无法实现这样赋值的
	//类似动态绑定,但是类型是cat,无法访问cat的方法
	fmt.Printf("%T \n", a)
	a.Say()
	c := a
	fmt.Printf("%T \n", c)

	//接收指针类型
	a = &cat{}             //这样的语法是没有错的
	fmt.Printf("%T \n", a) //类型也变为:*main.cat
}

type dog struct {
}

func (dog) Say() {
	fmt.Println("旺")

}
func (*dog) Eat() {
	fmt.Println("吃狗粮!")
}

// 当是接收者是指针时:接口只能接收指针类型
func InterfaceDemo2() {
	var a Animal = &dog{}
	//var a Animal=dog{} 这样是无法通过编译的
	a.Say()
}

type chick struct {
}

func (chick) Eat() {
	fmt.Println("虫子")
}

type bird struct {
	chick
}

func (bird) Say() {
	fmt.Println("我会飞!")
}

// 接口的方法不一定有一个类型A完全实现,一部分可以交给嵌入A的类型
func InterfaceDemo3() {
	var a Animal = bird{}
	a.Say()
}

// 接口之间可以嵌套组合成新的接口
type move interface {
	Animal
}

// 空接口:没有定义任何方法,也就是被任意类型实现
// 使用
// 可以给任意类型赋值,有时候可以作为函数参数,表明可以传入任意类型,
// 可以作为map的value部分,存放任意参数
// 但接口类型无法使用其实现类中独特方法和参数,当然我们可以通过类型断言来判断具体类型
// 当我们需要判断x变量是否是类型T时可以使用类型断言:
// t,ok:=x.(T) 返回一个T类型变量t,成功ok为true,失败false
func InterfaceDemo4(x interface{}) {

	var x2 interface{}
	x2 = "anc"
	v, ok := x2.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
	//注意:空接口使用广泛,但是空间的滥用只会使得代码变得更加抽象,同时更加会导致必要的内存浪费
}

func main() {
	InterfaceDemo1()
	InterfaceDemo2()
	InterfaceDemo3()
	InterfaceDemo4("hello")
}
