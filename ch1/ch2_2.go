package main

import "fmt"

// 变量声明
func main() {
	//一,变量的声明
	//1,ch1 变量名称 类型 =表达式  ==>go变量不赋值时,都有默认值
	var i int = 3
	//2,ch1 变量名称 =表达式  ==>go会自动推导变量类型
	var s = "hello world!"
	//3,go中可以按照组类声明变量
	var a, b int = 100, 0 //都是int类型
	fmt.Println("a=", a, "b=", b)
	a, b = b, a //变量值交换
	fmt.Println("a=", a, "b=", b)
	var x, y = "abc", true //自动推导
	fmt.Println(i, "\n", s, "\n", x, y)

	//二,简短变量的声明
	s2 := "简短变量声明1"
	fmt.Println(s2)
	i2, s3 := 2.2, "简短变量声明"
	fmt.Println(i2, s3)

	//三,指针变量,指针默认初始值为nil,go中任何变量拥有初始值从而都可以取出地址值
	var i3 int = 9
	var i3p *int = &i3        //获取i3的地址值
	fmt.Println("i3p=", i3p)  //输出地址值
	fmt.Println("*i3p", *i3p) //取出地址中存储的值
	*i3p++
	fmt.Println("i3p=", i3p)  //输出地址值
	fmt.Println("*i3p", *i3p) //取出地址中存储的值
	fmt.Println(i3)

	//new关键字的使用:创建变量的方式,new(T)返回T类型的指针
	var p2 *int = new(int)
	*p2 = 990
	fmt.Println(*p2)

}
