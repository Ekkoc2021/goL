package main

import "fmt"

// 指针:有别与c/c++中的指针,不可以进行运算
func main() {
	//go中空指针为nil
	var p *int //定义一个指向int类型的指针
	//*p = 100 不能够直接赋值
	//获取变量地址为指针赋值:&
	var i = 1
	p = &i
	*p = 2
	print(i) //2

	//new:func new(Type) *Type
	//返回改类型的指针,可以使用new函数初始化指针
	var p2 = new(int)
	*p2 = 100
	fmt.Println(*p2)

	//make:func make(t Type, size ...IntegerType) Type
	//make用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建,本身就是引用类型没有必要返回指针
	s := make([]int, 1)
	fmt.Println(s[0])

}
