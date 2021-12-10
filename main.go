package main

import (
	"fmt"
	"helloworld/internal"
)

func main() {

	// go 闭包的实现
	fmt.Println("go 闭包的实现")
	internal.Closures()

	// 递归的实现
	fmt.Println("递归的实现")
	internal.Recursion()

	// 指针的实现
	fmt.Println("指针的实现")
	internal.Pointers()

	// 接口实现
	fmt.Println("接口实现")
	internal.Interfaces()

	// 错误处理实现
	fmt.Println("错误处理实现")
	internal.ErrorImplement()

	// goroutine 的实现
	fmt.Println("goroutine 的实现")
	internal.Goroutine()

	// Channel 的实现
	fmt.Println("Channel 的实现")
	internal.Channels()

	// 通道缓冲
	fmt.Println("通道缓冲")
	internal.ChannelBuffering()

	// 通道同步
	fmt.Println("通道同步")
	internal.ChannelSynchronization()
}
