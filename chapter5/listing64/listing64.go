// 这个示例程序展示无法从另外一个包里
// 访问未公开的标识符
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing64/counters"
)

// main 是应用程序的入口
func main() {
	// 创建一个未公开的类型的变量
	// 这个类型用于保存告警计数
	counter := counters.alertCounter(10)

	// ./listing64.go:15: 不能引用未公开的名字
	//                                         counters.alertCounter
	// ./listing64.go:15: 未定义: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
