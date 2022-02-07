package main

import "fmt"

// 指针的实现：https://gobyexample-cn.github.io/pointers
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iprt *int) {
	*iprt = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
