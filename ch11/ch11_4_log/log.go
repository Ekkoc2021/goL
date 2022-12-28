package main

import (
	"fmt"
	"log"
)

// 打印日志
func demo1() {
	log.Println("Println打印的普通的日志。")
	log.Printf("Printf打印的%s日志。\n", "普通的")
	log.Fatalln("Fatalln打印的fatal的日志。") //在所有日志写入结束后调用在写入该日志然后调用os.Exit(1)
	log.Panicln("Panicln打印的panic的日志。") //打印完抛出异常
}

// 设置日志
func demo2() {
	//日志打印的具体内容
	//func Flags() int
	//设置日志具体内容
	//func SetFlags(flag int)

	//flog的选项
	/*const (
		// 控制输出日志信息的细节，不能控制输出的顺序和格式。
		// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
		Ldate         = 1 << iota     // 日期：2009/01/23
		Ltime                         // 时间：01:23:23
		Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
		Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
		Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
		LUTC                          // 使用UTC时间
		LstdFlags     = Ldate | Ltime // 标准logger的初始值
	)*/
	//使用多个可以使用+,|来搭配使用
	log.SetFlags(log.Llongfile | log.Lmicroseconds)
	fmt.Println(log.Flags())
	log.Println("Println打印的普通的日志。")

	//日志前缀设置
	log.SetPrefix("日志:")
	//设置日志输出位置:setOutput(w io.writer)
	//创建自己的日志对象:New(out io.writer,prifix string,flag int) *Logger

}
func main() {
	//demo1()
	demo2()
}
