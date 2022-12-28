package main

import (
	"flag"
	"fmt"
	"os"
)

// flag.Type()
// flag包用于解析启动时来自命令行的参数
// 获取命令行参数:go run xxx.go 参数....
func demo1() {
	//flag.NFlag() //返回使用的命令行参数个数
	//flag.Args() ////返回命令行参数后的其他参数，以[]string类型
	//flag.NArg() //返回命令行参数后的其他参数个数
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

// 命令行参数
func demo2() {
	//go run xxx.go -key value -key2 value ...
	//解析命令行中key-value同时赋值
	//同时当我们go run xxx.go --help时 会将用法默认值打印出来
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	//命令行参数解析
	flag.Parse()
	//使用参数
	fmt.Println(name)
	fmt.Println(age)
}

// flag.TypeVar()
func demo3() {
	var name string
	var age int
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.Parse()
}
func main() {

}
