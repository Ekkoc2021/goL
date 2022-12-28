package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Template:go语言模板有点像java的jsp:==>但是前后分离....,后面有空再回来看
// http:可以用来模拟浏览器发送请求,但...有空再看
// context:用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

/*
关于context详细解读的地址:https://www.cnblogs.com/qcrao-2018/p/11007503.html
Context接口:
type Context interface {
	返回截止的时间
	Deadline() (deadline time.Time, ok bool)

	返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，
	多次调用Done方法会返回同一个Channel
	Done() <-chan struct{}

	Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
	如果当前Context被取消就会返回Canceled错误；
	如果当前Context超时就会返回DeadlineExceeded错误；
	Err() error

	Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value
	并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；
	Value(key interface{}) interface{}
}

工作流程:多个协程共享一个上下文,上下文中封装同一个通道,通过一个通道的关闭与否去通知协程是否结束
取消函数:关闭通道
一个context可以允许有子context和父context,
当前context取消同时也会调用子context的取消函数,同时会解除和父context的关系

 Background()和TODO():都是返回一个context对象,源码中这两个对象实现接口的方法是空的,意味着使用这两个对象方法是没有意义的,而是将其做为根节点
Background作用于main函数,一般用做最顶层的context
基本上获取context的函数,必须传入一个context父对象,而一般是不建议传入空对象,这个时候可以同todo就可以被使用了(基本用不着)
大概的意思就是因为context之间的关系是树状的,你所有的context都必须从上面两个其中一个开始

context的分类
普通的context()==>Background()和TODO()这两个函数获取的context
可以被取消的context,取消意味通知结束:通过WithCancel(parent Context) (ctx Context, cancel CancelFunc)创建,返回一个context,一个取消函数(对象取消的方法是不对外开放的)
有截止日期的context,提前或者到时间自动取消:通过WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)创建
有时间限制的context,超时或者提前取消:通过WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)创建
可以存放k-v的context:通过WithValue(parent Context, key, val interface{}) Context创建,k必须是可以比较的,一个context只存放一对k-v
*/
