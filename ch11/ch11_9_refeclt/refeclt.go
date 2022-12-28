package main

import (
	"fmt"
	"reflect"
)

// 盲猜go框架必有反射,且是重要角色
/*
反射
一种在运行时更新和检查变量的值、调用变量的方法和变量支持的内在操作的机制，
但是在编译时并不知道这些变量的具体类型。

Go语言程序的反射系统无法获取到一个可执行文件空间中或者是一个包中的所有类型信息，
需要配合使用标准库中对应的词法、语法解析器和抽象语法树（AST）对源码进行扫描后获得这些信息。
*/

func demo1() {
	//无非就是关注变量的类型，变量的值，变量的方法
	//总共有四种大类型用于封装一个变量的方法：Type(类型)，Value,Filed,Metho
	//go中的类型是值go已经定义的类型,我们自己定义的属于结构体类型
	a := 2
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a)) //int 2

	//Type类型封装了包,名称,字段,方法
	//获取一个结构体的
	type Student struct {
		S string `tag:"你好世界" tag2:"hello world!"`
	}
	st := Student{}
	typ := reflect.TypeOf(st)
	//--------------------得到值操作
	value := reflect.ValueOf(st)
	fmt.Println(typ, value)
	//获取值的一些信息:类型,为其赋值,转为具体类型对象
	//kind方法返回的是基础的类型,type得到的可能是别名
	//可以将值返回一个接口,通过来下断言,获取这个变量然后使用
	//反射:可以反射类型的变量,然后可以得到接口类型变量,最后获得具体类型
	realV := value.Interface().(Student) //当然如果你知道具体类型,可以使用特定api返回类型
	fmt.Println(value.Type(), realV)
	//值的可写性:值的可写性是value类型的一个变量,判断是否能够修改值,基本都不可改
	//可取地址,和可写差不多,可以取地址表明可以改,指针value获取到的value都是可取地址的.对应取地址api:Addr() Value
	//slice传入的很明显传入的是指针,那么他对应slice[i]就是可改的
	fmt.Println(value.CanSet()) //判断是否可改
	//可写性很好理解,go中传递的变量如果不是指针,那么是值传递,值传递并不能影响到原来的值
	i2 := 3
	of2 := reflect.ValueOf(&i2) //of2是一个指针的值
	//我们只要通过这个指针获取到对应底层值,然后修改就能达到修改的修改
	elem := of2.Elem() //得到的value对象是可以修改的
	elem.SetInt(888)
	fmt.Println(i2)
	//修改结构体的值:
	//可修改的条件,结构体的字段开头大写(可被导出)
	//同时要传入结构体的指针
	type t3 struct {
		A int
		B string
	}
	t := t3{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
	//另外获取到的value对象提供了判断是否为空(空指针)或者无效(当查找到的字段为空时可以用这个)的方法

	//----------得到类型
	//获取该类型的一些信息
	fmt.Println(typ.Kind(),
		typ.String(),
		typ.Name(),
		typ.NumField(),
		typ.NumMethod(),
		typ.PkgPath())

	//获取指针指向的变量的类型
	p1 := new(map[string]interface{})
	tyof := reflect.TypeOf(p1)
	elemT := tyof.Elem() //tyof.Elem()用于获取指针指向的变量的类型
	fmt.Println(tyof, elemT)

	//获取类型的字段:可以通过索引,名称等获取一个类型的字段,还可以根据一个匹配函数去获取
	field := typ.Field(0)
	//获取字段上的信息:名称,类型,tag,Anonymous(是否是匿名字段)
	fmt.Println(field.Type, field.PkgPath, field.Name)
	//获取tag上的信息:tag的语法规则是相等严格的冒号前后不允许有空格`tag:"mess"`
	tag := field.Tag
	lookup, _ := tag.Lookup("tag2") //查询是否存在,存在就返回,不存在返回空,同时发货false
	fmt.Println(tag.Get("tag"), lookup)

	//----------创建实例
	//通过反射创建实例:New(Type)方法
	var a2 int
	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf(a2)
	// 根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA)
	// 输出Value的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind())

	//---------反射调用函数
	//函数对象的value可以使用call方法调用该函数,同时我们传入是一个value的切片作为函数的参数
	//通过反射调用函数
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	// 构造函数参数, 传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(103), reflect.ValueOf(20)}
	// 反射调用函数
	retList := funcValue.Call(paramList)
	// 获取第一个返回值, 取整数值
	fmt.Println(retList[0].Int())
	//通过调用方法也是如此,区分何时以参数类型传入对象

}

// 普通函数
func add(a, b int) int {
	return a + b
}
func main() {
	demo1()
}
