package main

import (
	"fmt"
)

// 结构体:结构体的内存是紧凑的
func main() {
	/*
	   结构体声明
	   1,如果类型相同成员定义可以写在同一行
	   2,匿名成员:声明类型但是不必声明其名称,好处就像类一样,直接继承其成员,可以直接访问
	   如果匿名成员之间出现名称冲突时还是需要通过名称来访问,名称就是该结构体的名称
	   结构体的可见性
	   1,通过首字母的大小写判定,大写公有,小写私有(可见性是对于其他包而言的)
	   type name struct{
	   	name type
	   	....
	   }
	*/
	type point struct {
		X, Y int
	}
	type circle struct {
		p point //匿名成员
		r int
	}
	type wheel struct {
		circle
		spokes int
	}

	//结构体实例化
	var w wheel
	w.p.X = 1 //非匿名成员访问还要多一层
	w.r = 10  //匿名直接访问
	fmt.Printf("%#v \n", w)
	//只有v表示只打印结构体的值,不打印名称,包含#打印名称,匿名成员也有名称且固定所以同类型的匿名成员只能有一个

	//通过取地址进行实例化:相当于new操作
	p2 := &point{} //结构体指针可以直接访问变量
	fmt.Printf("%T \n", p2)

	//实例化同时初始化
	p := point{1, 3} //定义变量同时初始化,必须初始化所有成员
	fmt.Println(p)
	var w2 wheel = wheel{ //键值对方式初始化
		circle: circle{
			p: point{
				X: 1,
				Y: 2},
			r: 2},
		spokes: 4} //包含匿名这样初始化好一点

	//结构体的比较性:成员可比较那么结构体可比较
	fmt.Println(w2 == w)

	//类型别名:对类型去别名,所有类型都适用
	type x = int //对int取别名,本质一样

	//匿名结构体
	var user struct {
		name     string
		id       int
		password string
	}
	user.password = "abc"

	//结构体的指针,go中的指针可以通过.直接访问成员
	p23 := new(point)
	p23.Y = 99
	fmt.Println(p23.Y)

	//结构体的方法
	student1 := student{"ekko", "0009"}
	name := student1.getName()
	fmt.Println(name) //ekko
	//go中所有的函数传参或者返回参数都是值传递
	//当数据量比较大时值的复制开销可能会比较大,这个时候需要考虑使用地址传递
	name = "jack"
	fmt.Println(student1.name) //ekko 并没有修改
	//结构体不存在构造函数,但是我们可以自定义函数使其充当一个构造函数

	//结构体赋值,go语言中的赋值操作基本都是copy值
	type d1 struct {
		name string
		id   int
	}
	d := d1{"hello", 33}
	d2 := d
	fmt.Println(d2)
	fmt.Println(&d, &d2) //地址不同,但是值相同

}

//结构体的方法定义
/* 接受者就是结构体的类型,返回参数为空可以省略
func (接受者变量 接受者类型) 方法名称(参数) 返回参数{
	函数体
}
*/
type student struct {
	name string
	id   string
}

func (s student) getName() string {
	return s.name
}
