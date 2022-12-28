package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//var data [16]byte
	//os.Stdin.Read(data[:])
	//os.Stdin.WriteString(string(data[:]))

	ReadDemo1()
	writeDemo()

	bufioWriteDemo()
	bufioReadDemo()

	ioutilWr()
	ioutilRe()
}

// 读取文件:读取到的数据是字节类型,可以使用read和readat读取,最后返回一个io.EOF的错误
func ReadDemo1() {
	f1, e := os.Open("./ch11/ch11_5_IO/hello.txt") //相对路径从当前项目的根目录
	defer f1.Close()
	if e != nil {
		fmt.Println(e)
	}
	var buff [10]byte
	data := make([]byte, 10) //缓存切片
	for {
		n, e := f1.Read(buff[:]) //返回读取到的字节个数,和一个错误
		if e == io.EOF {
			break
		}
		data = append(data, buff[:n]...) //展开
	}
	fmt.Println(string(data[:]))

}

// 写入数据:追加
func writeDemo() {
	/*0777表示：创建了一个普通文件，所有人拥有所有的读、写、执行权限
	0666表示：创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行
	0644表示：创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限，都没有执行权限*/
	f1, e := os.OpenFile("./ch11/ch11_5_IO/hello.txt", os.O_APPEND, 0666)
	defer f1.Close()
	if e != nil {
		fmt.Println(e)
	}

	f1.WriteString("你好")
	f1.Write([]byte("i am good!"))
	f1.Close()
}

// 带缓冲的读写:bufio包
func bufioWriteDemo() {
	f1, _ := os.OpenFile("./ch11/ch11_5_IO/hello.txt", os.O_CREATE|os.O_APPEND, 0666)
	defer f1.Close()
	write := bufio.NewWriter(f1)
	write.WriteString("写入了一条数asdfas据!\n") //os.O_WRONLY这种方式写入的数据是从开头写而且是替换
	write.WriteString("写入2了一条数asdfas据!\n")
	write.Flush()
}

func bufioReadDemo() {
	f1, _ := os.Open("./ch11/ch11_5_IO/hello.txt")
	defer f1.Close()
	reader := bufio.NewReader(f1)
	for true {
		d, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		}
		fmt.Println(string(d))
	}
}

// ioutil包的使用
func ioutilWr() {
	err := ioutil.WriteFile("./ch11/ch11_5_IO/hello2.txt", []byte("ioutil,你好世界\n 你好世界 !"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ioutilRe() {
	content, err := ioutil.ReadFile("./ch11/ch11_5_IO/hello2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
