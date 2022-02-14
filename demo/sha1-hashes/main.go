package main

import (
	"crypto/sha1"
	"fmt"
)

//SHA1 哈希: https://gobyexample-cn.github.io/sha1-hashes
/* md5 例子
    m := md5.New()
	m.Write([]byte(s))
	bs2 := m.Sum(nil)
	fmt.Printf("%x\n", bs2)
*/
func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
