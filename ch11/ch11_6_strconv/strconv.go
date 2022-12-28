package main

import (
	"fmt"
	"strconv"
)

// 常见函数:更多细节看官网
func main() {
	//strconv.Atoi():字符串转数字
	s := "89"
	i, _ := strconv.Atoi(s)
	i++
	fmt.Printf("%T", i)

	//strconv.Itoa():字符串到int
	s = strconv.Itoa(i)
	fmt.Printf("%T", s)

	//Parse系列:字符串转对应类型
	parseBool, _ := strconv.ParseBool("T") //1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	fmt.Printf("%T", parseBool)

	//整数
	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；

	//浮点数
	//func ParseFloat(s string, bitSize int) (f float64, err error)

	//FormatXxx系列:转换为对应格式的字符串并返回
}
