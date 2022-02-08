package main

import (
	"fmt"
)

//Recover: https://gobyexample-cn.github.io/recover
func mayPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}
