package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//结构体转为json或者json字符串转为结构体
	//tag:结构体标签使用``,tag是结构体的元信息,在运行的时候通过反射来读取
	//使用tag时一定要遵循`key:"vale"`的模式,tag语法容错很差,编译和运行都不会报错,但是会影响反射获取的数据

	type movie struct {
		Title string
		//key为json用于控制json包的编码和解码行为
		Year   int  `json:"released"`        //使用Tag,表示这个字段在json化时,字段名称变更
		Color  bool `json:"color,omitempty"` //omitempty选项表示当数据为空时不产生json对象
		Actors []string
	}

	m := movie{"hello", 2, false, make([]string, 2)}
	fmt.Println(m)

	//go语言中提供了将结构体转为json,yaml的字符串

	//结构体转json叫做编组(marshal):结构体字段能否编组取决是json包能否访问其字段,需要编排的字段应该大写
	//func Marshal(v any) ([]byte, error) v可以是切片或者是一个类型
	ms := make([]movie, 3)
	ms = append(ms, m)
	data, err := json.Marshal(ms) //当编组失败时,会返回一个错误
	if err != nil {
		log.Fatalf("json marshal failed :%s", err)
	}
	fmt.Printf("%s\n", data) //{"Title":"hello","released":2,"Actors":["",""]}]

	//marsha函数编排后数据可读性比较差,可以使用marshalindent函数
	//MarshalIndent(v any, prefix, indent string) ([]byte, error)
	//prefix每一个字段开始前缀,indent用于表示层级
	data2, err2 := json.MarshalIndent(m, "", "	")
	if err2 != nil {
		fmt.Println("编组失败!")
	}
	fmt.Printf("%s\n", data2)
	fmt.Println(string(data)) //[]byte转字符串
	//json==>结构体:func Unmarshal(data []byte, v any) error
	m2 := &movie{}
	e := json.Unmarshal(data2, m2)
	//传入指针类型,否则初始化失败,因为go中函数均是值传递
	if e != nil {
		fmt.Println("josn==>结构体 失败!")
	}
	fmt.Println(m2)

}
