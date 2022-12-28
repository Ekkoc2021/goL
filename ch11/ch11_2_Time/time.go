package main

import (
	"fmt"
	"time"
)

// 当前时间
func demo1() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//格式化:时间模板是go诞生的时间,其他都不行
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	//字符串解析为时间对象
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}

// 时间戳:Unix时间戳
func demo2() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("timestamp1:%v\n", timestamp1)
	fmt.Printf("timestamp2:%v\n", timestamp2)

	//时间戳转换为时间
	timeObj := time.Unix(timestamp1, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间操作
func demo3() {
	//time包对时间单位进行了定义,以纳秒为单位(你输入的是具体的数据请注意单位),详细看包内相关部分源码

	//add
	now := time.Now()
	fmt.Println(now.Add(time.Hour)) //加上1小时

	//sub
	fmt.Println(now.Add(time.Hour).Sub(now))

	//Equal
	fmt.Println(now.Add(time.Hour).Equal(now))

	//Before 是否在此之前: Before(u Time) bool
	//After 是否在此之后: After(u Time) bool

}

// 定时器
func demo4() {
	//Tick(d Duration) <-chan Time
	//定时,返回的是一个只允许取值的时间的通道,每隔d个单位的时间向管道输入当前时间
	tick := time.Tick(time.Second)
	times := 1
	for i := range tick {
		fmt.Println("你好,世界")
		fmt.Println(i)
		if times == 3 {
			break
		}
		times++
	}
}
func main() {
	fmt.Println("demo1----")
	demo1()
	fmt.Println("demo2----")
	demo2()
	fmt.Println("demo3----")
	demo3()
	fmt.Println("demo4----")
	demo4()
}
