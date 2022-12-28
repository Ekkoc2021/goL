package main

import "fmt"

func main() {
	forStatement()
}

// for语句:go中没有while语句,while能完成的for也能完成
func forStatement() {
	/*go中的for语句的语法和大多数语言基本一样 for init;condition;post{}
	 */

	//demo1
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	//demo2 init和condition可以不用写
	for ; ; sum++ {
		fmt.Println(sum)
		if sum > 5051 {
			break
		}
	}

	//demo3 只包含condition,前后的分号可以省略
	for sum <= 5053 {
		fmt.Println(sum)
		sum++
	}

	//demo4:下面两种无限循环都是等价的
	for true {
		break
	}
	for {
		break
	}

	//demo5 可以通过for进行迭代操作,像极了python(≧∇≦)ﾉ
	//迭代对象:字符串,切片,数组,集合,channel(只有value没有key值)
	m := map[string]string{"1": "jack", "2": "mike"}

	//可以完全不获取k,v值
	for range m {

	}
	//值获取k
	for k, _ := range m {
		fmt.Printf("%s \n", k)
	}
	//只获取v
	for v := range m {
		fmt.Printf("%s \n", v)
	}
	//k,v都获取
	for k, v := range m {
		fmt.Printf("%s:%s\n", k, v)
	}
	//注意:range迭代获取到的k,v是一个赋值过程不是原先的对象,这个很好理解,go中全是值传递,除非是指针
	for k, v := range m {
		v = "hello"
		fmt.Printf("%s:%s 修改后的v=%s\n", k, m[k], v)
	}
	//range迭代的对象是原对象的副本,对原对象的修改不会影响range
	//如果你也会python可以去试试python的有影响吗?
	arr := [2]int{1, 2}
	for _, v := range arr {
		arr[1] = 100   //在打印v之前,我们修改数组的值
		fmt.Println(v) //第二次获取到的v值依旧是原来的,没有改动的
	}

}
