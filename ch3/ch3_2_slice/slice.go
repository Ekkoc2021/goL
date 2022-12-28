package main

import (
	"fmt"
)

func main() {
	//Go中的slice依赖于数组，它的底层就是数组，所以数组具有的优点，slice都有。
	//且slice支持可以通过append向slice中追加元素，长度不够时会动态扩展，
	//通过再次slice切片，可以得到得到更小的slice结构，可以迭代、遍历等。
	//切片:将数组切片后返回对应指针

	strs := [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	//取出数组的切片:会返回一个容量为原来数组长度,长度为切割的长度的切片
	var s1 = strs[1:3]   //切出[0,2)
	fmt.Println(cap(s1)) //容量为9,why?
	fmt.Println(len(s1)) //长度为2
	fmt.Println(s1)
	fmt.Println(s1[1])
	s1 = append(s1, "abc") //切片扩容
	fmt.Println(s1, strs)  //[b c abc] [a b c abc e f g h i j]
	//切片指针指向的是原来的数组并非重新分配内存,扩容会在原来数组上做修改,如果超过数组长度会如何扩容?
	fmt.Printf("%T \n", s1) //[]string
	//还可以这样切
	a1 := strs[:6:10] //表示容量为10,超过(10-0)扩容,10应该小于原来数组的长度
	fmt.Println(a1)

	s1[0] = "x"
	fmt.Println(strs[0]) //数据改变了,说明slice底层维护的是指针,修改会变动原来数组的值

	//定义以切片变量
	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2)
	//slice之间是无法比较的,也就是无法使用!=,==
	//slice与数组之间的区别,slice底层维护一个数组变量,slice的指针指向数组,唯一合法的比较是判断slice是否为空nil

	//内置的make,可以创建指定初始长度和指定初始容量的的slice make([]T,len,cap) []T
	s3 := make([]int, 2) //容量和长度相同
	fmt.Println(s3)
	s4 := make([]int, 2, 3) //长度应该小于容量,这很好理解
	fmt.Println(s4)

	//追加:append([]T,...args) []T
	//向切片中追加一些数据,可以多个

	s5 := []int{1, 2, 3}
	fmt.Println(cap(s5)) //3
	s5 = append(s5, 9)
	s5 = append(s5, 9)
	s5 = append(s5, 9)
	s5 = append(s5, 9)
	fmt.Println(cap(s5)) //12
	fmt.Println(s5)

	s4 = append(s4, 8)
	s4 = append(s4, 8)
	fmt.Println(cap(s4))

	//关于rune类型的一个demo
	var runes []rune
	for _, r := range "hello,世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q", runes) //输出结果:['h' 'e' 'l' 'l' 'o' ',' '世' '界'],注意不是纯数字

	//append每次都返回一个对应类型的切片,当不需要扩容时容量是不变的,当需要扩容时,容量是改变的

	s6 := make([]int, 5, 6) //设置一个长度为1,初始化长度为6的切片
	fmt.Println(cap(s6))
	z := s6[:3]
	fmt.Println(cap(z)) //6,类似于切割数组

	//返回的切片显然和原来的是不相等的
	s7 := make([]int, 1, 3)
	s8 := append(s7, 2)
	fmt.Println(&s7 == &s8) //false

	//当容量不够时,扩容机制是翻倍的,当超过1024时扩容倍数是1.25而不是2倍
	fmt.Println(cap(s8)) //3
	s8 = append(s8, 2)
	fmt.Println(cap(s8)) //3
	s8 = append(s8, 2)
	fmt.Println(cap(s8)) //6

	/*如果超过容量则会新建一个数组，新旧切片指向不同的数组；
	如果没有超过容量则在原有数组上追加元素，新旧切片指向相同的数组，
	这时对其中一个切片的修改会同时影响到另一个切片。*/
	x := make([]int, 2, 3)
	x[0] = 1
	x[1] = 2
	x1 := append(x, 90)
	x1[0] = 98         //x也改变了
	fmt.Println(x, x1) //[98 2] [98 2 90]
	x2 := append(x, 990)
	fmt.Println(x, x1, x2) //[98 2] [98 2 990] [98 2 990] x1发生改变了
	x3 := append(x2, 1000)
	x3[0] = 0
	fmt.Println(x, x1, x2, x3) //[98 2] [98 2 990] [98 2 990] [0 2 990 1000]
	//切片之间共享相同数组,扩容后不再共享相同数组
	//如果我们直接从数组进行切片,那么切片指针会指向该数组,容量会变为:数组长度-(切片开始位置+1)
	//其实就是切片开始位置到数组末尾的长度
	arr := [3]int{100, 99, 98}
	x4 := arr[2:3]
	fmt.Println(cap(x4))

	//复制函数:copy(dst []T,src []T) int
	//作用:将源切片的数据复制到中目标切片中
	//返回值:复制成功的个数,显然不用担心超出容量的问题,源切片小于目标切片那么成功值为源切片的长度
	s9 := []int{1, 2, 3, 4, 5}
	s10 := make([]int, 2)
	fmt.Println(copy(s10, s9)) //2
	fmt.Println(s10)           //1,2

	//copy函数和直接切片的区别:copy函数只是值,而切片是直接指针指向原地址
	s11 := s9[:2]
	fmt.Println(&s11[0] == &s9[0]) //ture
	fmt.Println(&s10[0] == &s9[0]) //false
	s12 := s9                      //和数组不同切片本身值传递
	s12[0] = 900
	fmt.Println(s12, s9)
	fmt.Println(&s12 == &s9) //false,地址不同,依然是值传递,不过底层维护的是地址值

	//多维的切片:与数组有区别
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data)

	//通常函数传参数,而且参数是数组,那么可以使用切片,切片不定长,数组定长且不是地址传递,如果传入数组地址那么就是定长的不够灵活

}

// string是不可改的,但是可以通过slice切片后修改string
func sliceToString() {
	str := "Hello world"
	s := []byte(str) //中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
}
