package main

import "os"

//Panic: https://gobyexample-cn.github.io/panic
func main() {
	// panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
