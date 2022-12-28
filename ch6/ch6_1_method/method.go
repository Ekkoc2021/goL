package main

import "fmt"

// go语言中的方法
// go语言中可以给结构体类型定义方法
type student struct {
	id int
}

func (s student) printS() {
	fmt.Println(s)
}

// 无论传入的是指针还是变量,效果相同
func (s *student) printS2() {
	fmt.Println(s)
}

type teacher struct {
	student
}
type class struct {
	*student
}

//go语言中的方法也是一个类型
//1,当赋值的是实例化的方法,那么可以直接使用
//2,赋值的表示实例化的方法,需要传入对象才可以调用,对象位置为第一个
//(也很好理解,方法引用实例化对象,不传入会导致错误)==>此时就成了普通函数

type T1 struct {
	id int
}

func (t T1) T1Method1() {
	fmt.Println(t.id)
	t.id = 99999
}
func MethodDemo1() {
	t := T1{100}
	f := t.T1Method1
	f() //默认传入了t,但是只是值赋值,其实和简单函数差不多,只不过多了一条规则
	fmt.Println(t.id)
	f2 := T1.T1Method1
	f2(t)
}
func main() {
	s := student{2}
	s.printS()
	s.printS2()

	s2 := &student{}
	s2.id = 0
	s2.printS2()

	t := teacher{}
	t.id = 9
	t.printS()
	t.printS2()

	c := class{}
	//c.printS()出现了空指针
	c.student = &student{}
	c.printS()

	MethodDemo1()

}
