package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//行过滤器: https://gobyexample-cn.github.io/line-filters
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
