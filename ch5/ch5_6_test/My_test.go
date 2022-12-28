package ch5_6_test_test

import (
	"goL/ch5/ch5_6_test/mycode"
	"testing"
)

// go中测试文件名称必须以_test.go结尾,go build 不会打包这类结尾的文件
// go中单元测试的函数必须以Test开头,且传入固定的对象
func TestAdd(t *testing.T) { //我们测试mycode中的一个函数Add2函数,不再需要使用mian方法
	test.Add2(1, 2)
}

//go的单元测试类型还有很多,功能也很强大
//优秀的程序员有编写单元测试的习惯
