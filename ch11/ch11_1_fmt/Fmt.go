package main

import (
	"bufio"
	"fmt"
	"os"
)

// 输入与输出函数
func main() {
	//将内容自动转入标准输出
	fmt.Println("hello")       //换行
	fmt.Print("world!")        //不换行
	fmt.Printf("你好!%s", "世界.") //格式化输出
	//将内容输入对应输出流
	fmt.Fprint(os.Stdout, "你好!") //os.Stdout=标准输出流
	fmt.Fprintln(os.Stdout, "你好!")
	fmt.Fprintf(os.Stdout, "你好!%s", "世界")
	//将内容转换为字符串然后返回
	/*
		Sprint(a ...any) string
		Sprintln(a ...any) string
		Sprintf(format string, a ...any) string
	*/
	//内容信息包装成一个错误对象上并返回
	err := fmt.Errorf("错误")
	fmt.Println(err)

	//格式化占位符:查资料
	//%T打印类型
	//%d打印整数
	//%f打印浮点数
	//%s打印字符串

	//获取输出:Scan Scanf Scanln是不能连用的
	//Scan 读取到对应变量中,空格分割
	//var s string
	//var d int
	//fmt.Scan(&s, &d) //会返回赋值成功的数量,读取信息错误会返回一个错误
	//fmt.Println(d, s)
	//fmt.Scanln("读取缓存数据", &s, &d)
	//fmt.Println(s)
	////需要按照给出的格式进行录入,不同数据间的格式需要用空白分割
	//scanf, err := fmt.Scanf("你好%s %d", &s, &d)
	//fmt.Println(scanf, err)
	//fmt.Println(d)
	//fmt.Println(s)

	//bufio.NewReader 有很多方法,看文档
	reader := bufio.NewReader(os.Stdin)
	//ReadString(delim byte) (string, error) :读取字符串,delim表示分界
	readString, _ := reader.ReadString('\n')
	fmt.Println(readString)

	//Fscan Fscanln Fscanf 从io.reader中读取数据,基本和scan系列用法一致
	var (
		s3 string
		i1 int
	)
	//解决了无法scan系列中无法连续读取的问题
	fmt.Fscan(os.Stdin, &s3, &i1) //换行符读取到数据就结束 不读取缓冲里面的换行符号
	fmt.Println(s3, i1)
	fmt.Fscanln(os.Stdin, &s3, &i1) //这里会读取到缓冲里面的换行符号
	fmt.Println(s3, i1)
	fsca, err := fmt.Fscanf(os.Stdin, "输入s3%s 输入i1%d", &s3, &i1) //这里会读取到缓冲里面的换行符号
	fmt.Println(fsca, err)
	fmt.Println(s3, i1)
}
