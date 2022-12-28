package main

import (
	"fmt"
	"time"
)

//go语言支持递归

// 斐波那契数列
func fibonaci1(n int) (result int) {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonaci1(n-1) + fibonaci1(n-2)
}

// 上面那种做法性能其实是有浪费的:fibonaci1(n-1)进行递归时,必然用到fibonaci1(n-2)的结果但是这里是直接再次调用
func fi(begin, end, n1, n2 *int) {
	if *end == *begin {
		return
	}
	*n2 = *n2 + *n1
	*n1 = *n2 - *n1
	*begin += 1
	fi(begin, end, n1, n2)
}
func fibonaci2(n int) (result int) {
	begain := 1
	n1 := 0
	n2 := 1
	fi(&begain, &n, &n1, &n2)
	return n2
}
func main() {
	now := time.Now()
	start1 := now.UnixNano()
	fmt.Println(fibonaci1(30)) //34
	now = time.Now()
	end1 := now.UnixNano()
	fmt.Println(end1 - start1) //2548400

	now = time.Now()
	start2 := now.UnixNano()
	fmt.Println(fibonaci2(6000))
	now = time.Now()
	end2 := now.UnixNano()
	fmt.Println(end2 - start2) //520600

}
