package main

import "fmt"

// 城市常量定义，使用 iota 生成相关的常量值
// iota 是 Go 语言中的一个关键字，用于在常量定义中自动递增数值
const (
	BEIJING   = 10 * iota // BEIJING 表示北京，值为 10
	SHANGHAI              // SHANGHAI 表示上海，值为 20
	SHENZHEN              // SHENZHEN 表示深圳，值为 30
	GUANGZHOU             // GUANGZHOU 表示广州，值为 40
)

// main 函数是程序的入口点
func main() {
	// 定义一个名为 length 的整型常量，值为 1024
	const length int = 1024
	// 打印 length 常量的值
	fmt.Printf("length is %d\n", length)
	// 打印 BEIJING 常量的值
	fmt.Println("BEIJING = ", BEIJING)
}
