package main

import (
	"fmt"
)

func main() {
	//定义未赋值
	var arrays [2]int
	fmt.Println(arrays[0])

	//赋值
	var arr2 [2]int = [2]int{4, 2}
	for _, a := range arr2 { //第一位获取到索引值,第二位为具体值,索引值不使用,使用_代替
		fmt.Println(a)
	}

	//根据所给的数的个数去判断数组长度:使用...表示未知长度
	var arr3 = [...]int{5, 6, 7, 8}
	for _, a := range arr3 {
		fmt.Println(a)
	}

	//数组可以指定索引值
	var strarr = [...]string{3: "hello", 2: "world"}
	fmt.Println(strarr[0])
	fmt.Println(len(strarr)) //4 指定索引值后,自动推断的长度,会根据索引最大值来确定

	//数组类型是可以比较的:!=,== 会根据元素个数,索引位置的值判断是否相等,当然得要求是同类型
	fmt.Println(arrays == arr2)
	fmt.Println(arrays) //go语言中的数组是允许打印的

	//go中数组是值类型,赋值和传参会复制数组,属于值拷贝,建议使用数组指针或者切片
	arrx := [2]int{199, 200}
	arrx2 := arrx
	arrx2[0] = 188
	fmt.Println(arrx, arrx2) //[199 200] [188 200]

	//go中的多维数组
	b := [...][...]int{{1, 1}, {2, 2}, {3, 3}, {2}}
	fmt.Println(b)

}
