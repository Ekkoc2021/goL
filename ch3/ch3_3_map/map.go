package main

import (
	"fmt"
)

func main() {
	//map的底层是哈希表的维护,默认值是nil
	//map类型的定义:map[KeyType]ValueType
	m1 := map[string]string{"url": "mysql", "user": "root", "password": "1234"} //定义一个map
	//取值
	var i string = m1["user"]
	fmt.Println(i) //root
	//go中map的取值是安全的,当是基本数据类型时没有时会返回默认值
	m2 := map[string]int{"1": 2}
	v, e := m2["url2"] //v表示取出的值
	fmt.Println(v, e)  //mysql true
	if _, ex := m2["url2"]; !ex {
		fmt.Println("不存在url2")
	}
	//go中map是可以进行运算
	re := m2["1"] + 1
	fmt.Println(re)
	//go中map不可以取值的:因为map是会进行扩容的,扩容后的地址无效:&m2["1"]是错误的语法
	//map的迭代是完全随机的,每次都不同,故意如此.但是可以通过区key后进行排序完成有序遍历
	for k, va := range m1 {
		fmt.Println(k, va)
	}
	//go中不提供set,可以用map实现一个set,因为map的key本身就是不相同的

	//创建map方式:
	//如前面的例子声明的同时指定,或者不指定
	//使用make:make(map[KeyType]ValueType, [cap])
	m3 := make(map[string]string, 1) //容量不是必须指定
	m3["a"] = "hello"
	m3["b"] = "world"
	for k1, v1 := range m3 {
		fmt.Println(k1, v1)
	}

	//删除键值使用delete函数, delete(map, key)
	delete(m3, "a")
	delete(m3, "ac") //delete函数是安全的,不存在ac键也不会不会报错
	for k1, v1 := range m3 {
		fmt.Println(k1, v1)
	}

	//map的底层维护的是一个数组的hash表,通过hash函数映射到数组的位置
	//扩容意味着丢弃原先的地址,所以前面的取地址值是非法的
	//数组下标存储一个bucket,每个bucket可以存储8个kv

}
