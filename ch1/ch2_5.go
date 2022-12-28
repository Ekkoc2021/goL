package main

// 包和文件
// 一个包的代码可以保存在一个或者多个.go文件中,要求包名称一致
//一般一个目录下的多个文件的包名称一致,且和目录名称一致,导入才不会出错
//可被导入的包应该和项目都放在gopath的src目录下,可以拥有多个gopath
//包中的开头字母是大写的才可被其他包导入后使用

import (
	"fmt"
	"goL/ch1/packageTest"
)

func main() {
	var i packageTest.Id = 1
	fmt.Println(i)
}
