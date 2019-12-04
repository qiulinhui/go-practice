// 这个示例程序展示不是总能
// 获取值的地址
package main

import "fmt"

// duration 是一个基于 int 类型的类型
type duration int

// 使用更可读的方式格式化 duration 值
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

// main 是应用程序的入口
func main() {

	// duration(42).pretty()
	// 修改后
	d := duration(42)
	d.pretty()

	// ./listing46.go:17: 不能通过指针调用 duration(42) 的方法
	// ./listing46.go:17: 不能获取 duration(42) 的地址
}
