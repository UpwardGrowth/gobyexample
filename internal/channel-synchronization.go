package internal

import (
	"fmt"
	"time"
)

// 通道同步：https://gobyexample-cn.github.io/channel-synchronization
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ChannelSynchronization() {
	done := make(chan bool)
	go worker(done)

	<-done
}
