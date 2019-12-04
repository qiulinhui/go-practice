// Package entities 包包含系统中
// 与人有关的类型
package entities

// user 在程序里定义里定义一个用户类型
type user struct {
	Name  string
	Email string
}

// Admin 在程序里定义了管理员
type Admin struct {
	user   // 嵌入类型是未公开的
	Rights int
}
