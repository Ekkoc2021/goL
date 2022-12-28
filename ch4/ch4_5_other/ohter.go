package main

import (
	"fmt"
)

func main() {
	gotoAndBreakAndContinue()
}

// goto,break,continue
func gotoAndBreakAndContinue() {
	//和大多数语言一样配合lable使用

	//goto:跳到指定标签的位置
	if i := 1; i == 1 {
		goto a1
	}
	fmt.Println("我被跳过了!")
a1:
	fmt.Println("跳过了上一句!")
	//break,默认跳出当前的for,select,switch语句,配合标签可以明确跳出具体的语句
	//continue,跳过本次循环直接进入下一次,默认为当前的循环,配合标签可以具体指定循环
}
