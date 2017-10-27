package main

import (
	"fmt"
	"unicode"
)

// 练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回

func RmSpace(b []byte) []byte {
start:
	n := len(b)
	if n == 1 {
		return b
	}
	for i := 0; i < n-1; i++ {
		if unicode.IsSpace(rune(b[i])) && b[i] == b[i+1] {
			copy(b[i:], b[i+1:])
			b = b[:n-1]
			goto start
		}
	}
	return b
}

func main() {
	b := []byte("fdsa  gsag  grsag     ghrsa g     gers")
	fmt.Println(string(b))
	fmt.Println(string(RmSpace(b)))
}
