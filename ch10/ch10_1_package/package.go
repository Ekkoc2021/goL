package main

import io "fmt"
import _ "os" //匿名导入

// 导包:单个包可以不用小括号,多个必须用
// 导入的包如果用不到需要在此之前添加下滑写表示匿名导入,并不使用
// 导入的包如果有同名,必须使用别名区分
// 内部包,go中路径中包含internal被称为内部包只能被其父母了的包访问:包/a/b/internal/c只能被a/b/*访问
func main() {
	io.Println()
}
