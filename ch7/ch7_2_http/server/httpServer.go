package main

import (
	"fmt"
	"net/http"
)

// 使用go快速搭建一个简单的http服务器
func main() {
	http.HandleFunc("/go", sayhello) //当访问/go路径时,调用回调函数sayhello
	http.HandleFunc("/", welcome)    //跳转欢迎页面
	http.ListenAndServe("localhost:8080", nil)
	//handler是一个接口,有一个方法被实现,我们可以通过这个接口的的方法实现请求分发,这里我们不作分发
}
func sayhello(w http.ResponseWriter, r *http.Request) { //要传入指针,确保两个request相同
	//request中封装了请求信息,可以通过对应方法实现调用
	fmt.Println("/go 被访问了")
	w.Write([]byte("<h1>收到!</h1>"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>感谢访问,我的网站!</h1>"))
}
