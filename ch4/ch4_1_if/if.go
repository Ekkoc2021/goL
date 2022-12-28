package main

import "fmt"

func main() {
	ifStatement()
}

// if语句
func ifStatement() {
	//if语句:go语言的if语句结构和其他语句的结构是相似的只不过语法上会有区别
	/*
		if 代码块1 {
			代码块2
		}
		代码块1将整个if语句内的所有代码块包含在一起,这意味着if内的代码块可以访问1中定义的
		go圣经中称为词法域,而代码块1称为隐式词法域
		代码块1只允许这样的格式:语句;条件
		语句可以是定义变量,调用函数,只要能够在一条语句完成的都行,但是像这种是非法的:var i int=2
	*/
	//demo1
	var a int = 2
	if fmt.Println("dmeo1"); a < 5 {
		fmt.Println("a小于5")
	}

	//demo2
	if a, b := 3, 5; a < b {
		fmt.Println(a)
	}

	//demo3
	if c := "helloworld"; a < 1 {
		fmt.Println("a小于1")
	} else {
		fmt.Println("a>1" + c)
	}

	//demo4
	if c := "helloworld"; a < 1 {
		fmt.Println("a小于1")
	} else if b := "dlrowolleh"; a < 2 {
		fmt.Println(c)
	} else {
		fmt.Println(b)

	}
}
