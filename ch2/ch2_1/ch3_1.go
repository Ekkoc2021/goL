package main

import (
	"fmt"
)

// 整型,浮点型,复数,布尔型,字符串,常量,二元运算符号
func main() {
	dataType()
	constant()
	operator()

}
func dataType() {
	//应该注意一点,就是各种数据类型之间是不允许强转,
	integer() //整形
	float()
	complax()
	boolean()
	String()
}

// 1,整型
func integer() {
	// 整形:int int8 int16 int32 int64 byte byte等价于int8
	fmt.Println("\n数据类型:")
	var i int8 = 127 //8位整形
	fmt.Println(i + 1)
	var i2 int = int(i) //int类型与操作系统相关,而且类型之间必须存在显式转换
	fmt.Println(i2 + 1)
	var i3 uint8 = 134 //无符号8位
	fmt.Println(i3)
	var i4 int64 = 8989 //有符号64位
	fmt.Println(i4)

	//进制:0开头8进制,0x或者0X开头16进制,
}

// 2,浮点型
func float() {
	//浮点数:float32 float64 ==>没有float
	fmt.Println("\n浮点数:")
	var f float32 = 3.141592654 //32位输出:3.1415927
	fmt.Println(f)
	var f2 float64 = 3.141592653545555666 //输出3.1415926535455556
	fmt.Println(f2)

	//细节1:纯小数,小数点前面的如果是0可以省略
	var f3 float64 = .32
	fmt.Printf("纯小数,省略小数点前面的0:%f \n", f3)
	//细节2:可以使用科学计数法
	var f4 float32 = 1.23e2 //计算机并不能精确的表示所有的浮点数
	fmt.Printf("科学计数法:1.23e2=%f \n", f4)
}

// 3,复数
func complax() {
	// 复数类型:complex64(float32),complex128(float64)
	// 本质是通过浮点数组合构建的组合类型
	fmt.Println("\n复数:")
	var x complex64 = complex(3.141592654, 3.141592654)
	fmt.Println(x) //(3.1415927+3.1415927i)
	var x2 complex128 = complex(3.141592654, 3.141592654)
	fmt.Println(x2) //(3.141592654+3.141592654i)

	//一些常用的复数操作函数
	var x3 = complex128(x) //转型
	fmt.Println(x3 * x2)
	fmt.Println(real(x3 * x2)) //实数
	fmt.Println(imag(x3 * x2)) //虚数

	//自然方式书写
	var x4 complex64 = 2 + 3i
	fmt.Println(x4)
	//复数是可以使用==,!=做比较的
	//math/complx包提供了很多处理复数的运算
}

// 4,布尔:默认是false
func boolean() {
	println("\n布尔类型:")
	var b1 bool
	fmt.Println(b1)
	b1 = true
	fmt.Println(b1)
	//布尔类型不会隐式转换为0或者1,需要配合使用if语句进行转换

}

// 5,常量
func constant() {
	println("\n常量:")
	//常量:编译期就完成了常量的运算,编译期就确定了值,从而可以用来指定数组的长度
	const r = 9
	const ( //组合声明多个常量
		e1 = 2.718
		pi = 3.14
	)

	//常量计数器iota,类似枚举
	const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
	const (
		n12 = iota //0
		n22        //1
		_          //丢弃该值，常用在错误处理中
		n42        //3
	)
	//插队
	const (
		n13 = iota //0
		n23 = 100  //100 本来该赋予1,用100修改赋予的1
		n33 = iota //2
		n43        //3
	)
	const n5 = iota //0
	const (
		_  = 1 << (10 * iota)
		KB // <<移位操作，速度比乘除法快 相当于1<<10
		MB //表示MB的规则和上面一致==>等价 MB= 1 << (10 * iota)
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	println(MB)
	println(1 << 20)
	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)

	//许多常量没有一个明确的基础类型,ZB之后的常量已经超出了go整型的表示范围,编译器为这些常量准备了更高精度的算术运算
	//fmt.Printf("%T \n", ZB) //不能打印但是可以运算
	fmt.Println(YB / EB)

	//我们知道int类型和int8,int16等这些是不能直接赋值的,需要转换
	//var itest int =8
	//var i64test int64=itest

	fmt.Printf("%T \n", PB) //int类型
	var i64 int64 = PB      //编译通过
	fmt.Println(i64)

	//无类型的常量赋值给有类型的变量,在精度不丢失的情况下会隐式转换类型
	var f2 float64 = 2 + 0i
	fmt.Println(f2)

	//var f2t float64 =2.0
	//var f4 int =f2t //这样是不被允许的的
	//但是常量却可以
	const test = 2.0
	fmt.Printf("%T \n", test) //float64
	var f3 int = test         //float是不能赋予int的,如果2.0改为2.2将非法
	fmt.Println(f3)

	//自动推导类型时,float和complex自动推导都是最高为,也就是float64,complex128
	//但是int类型默认推导为int,也就是精度依旧不确定,或者说由操作系统决定
	ix := 10
	fmt.Printf("%T \n", ix) //输出int
	//想用获取确定的精度,使用对应类型进行强转,或者声明时指定

}

// 6,字符串类型:可以看到go取消了字符类型,其实字符类型完全可以融合在字符串类型,这一点很赞
func String() { //go中的文件总是utf-8编码,不用再担心中文问题
	fmt.Println("\n字符串类型:")
	var s1 string
	fmt.Println(s1) //""
	s1 = "hello world"
	fmt.Println(s1) //"hello world"

	//字符串常用的函数
	fmt.Println(s1[0]) //索引取值,但是取值是ASCII码,而不是字符,可以直接强转为string
	fmt.Println(s1[:]) //和py字符串很像这点对于有py基础的来说很快乐
	fmt.Println(len(s1))
	fmt.Println("goodbye" + s1[5:])
	//字符串是不可改的:s1[0]='x'这样是非法的

	var s2 rune = 'c'
	println(s2) //99
	var s3 = "b"
	println(s3 + string(s2))
	fmt.Printf("s2类型:%T s1类型:%T \n ", s2, s1) //可以看到s2是int32类型,而s1是string

	//Unicode对应go中rune类型等价于int32
	//4个标准库:bytes,strings,strconv,unicode,查看详细文档包含使用实例

	//字符串的多行输出:使用反引号,tab键上面
	s := `
		hello
		world
    `
	fmt.Println(s)

	//字符串常见操作
	s = fmt.Sprintf("你好,%i", 2)
	fmt.Println(s)
}

// 7,二元运算符号
func operator() {
	fmt.Println("\n二元运算符号:")
	//go中运算符号优先级 优先级降序由上到下
	//* / % << >> & &^
	//+ - | ^
	//== != < <= > >=
	//&&
	//||

	//&^ :左边与右边不同则取左边,相同取0
	fmt.Println(1 &^ 0)
	fmt.Println(1 &^ 1)
	fmt.Println(0 &^ 1)

	//^ : 异或,相同位0,不同为1
}
