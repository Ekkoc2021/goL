package main

import "fmt"

// 类型:一般出现在包的一级,从而外包也能使用
type Stu_id int

// 声明了一个类型,实际就把int进行包装了一下,声明的数据类型运算方式和底层类型一致
// 类型可以包含方法的,第6张会详细讲解,比如下面这个例子,将类型格式化后返回类似于java的toString方法
func (id Stu_id) String() string {
	return fmt.Sprintf("%g", id)
}

// 注意:如果类型首个字母是大写的则在外部包也能使用,而中文等等被看做小写字母(至少在go2之前是这样的)
func main() {
	var i Stu_id = 10
	var i2 int = 9
	fmt.Println(i)
	fmt.Println(i2)
	//i=i2//报错
	i = Stu_id(i2) //类型转换 如果是指针类型的话:(*int)(i2)
}
