package main

import "fmt"

func main() {
	switchStatement()
}

// switch语句
func switchStatement() {
	//demo1
	var a int = 1
	switch fmt.Println("hello"); a {
	case 1:
		fmt.Println("a=1")
		//break,可以不用break,go中Switch是不会出现case穿透现象的
		//但是我们可以通过 fallthrough 来实现穿透
		fallthrough
	case 2:
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	}

	//demo2
	switch a := 2; a {
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println(a)
	}

	//demo3:Switch可以用来判断类型
	var x interface{}
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case rune:
		fmt.Println("rune")
	default:
		fmt.Printf("%T\n", x)
	}

	//demo4 可以忽略Switch后面表达式,case后面表达式可以用true或者false来判断是否进入对应代码块
	a = 1
	switch {
	case a < 1:
		fmt.Println("a<1")
	case a < 2:
		fmt.Println("a<2")
	}
}
