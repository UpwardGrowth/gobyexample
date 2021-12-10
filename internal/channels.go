package internal

import (
	"fmt"
)

// channels：https://gobyexample-cn.github.io/channels
func Channels() {
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
}
