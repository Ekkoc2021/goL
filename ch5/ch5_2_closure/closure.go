package main

import (
	"fmt"
)

// 闭包:由函数及其相关引用环境组合成的实体
// go语言的闭包函数,当一个函数A返回值是另外一个函数B时,函数B可以直接访问在A中定义的变量
// B中访问的A中定义好的变量称为捕获变量,B称为闭包函数
func main() {
	A := func() []func() {
		b := make([]func(), 2)
		for i := 0; i < 2; i++ {
			b[i] = func() {
				fmt.Println(i)
			}
		}
		return b
	}
	b := A()
	b[0]() //2
	b[1]() //2
	//都会输出2,因为两个函数访问的是同一块地址,访问时该地址的值为2
}
