package internal

import "fmt"

// 通道缓冲：https://gobyexample-cn.github.io/channel-buffering
func ChannelBuffering() {

	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
