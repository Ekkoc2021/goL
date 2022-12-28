package main

import (
	"fmt"
)

// 变量的赋值
func main() {
	//赋值
	var i int = 10 + 2
	i = 3
	j := 9
	j, i = 10, 90 //j=10,i=90
	//f,err=os.Open("foo.txt") //右边产生了多个值
	//当产生的值不使用时可以使用空白符号:_ 取代,当没有产生值时会根据条件报异常

	//自增加,不支持--j,--i,j=i--
	j++ //自增
	i-- //自减

	fmt.Println(i + j)

}
