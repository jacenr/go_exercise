package main

import (
	"fmt"
)

func main() {
	a := make([]byte, 1, 5)
	b := a[:len(a)+2]
	fmt.Println(a)
	fmt.Println(b)
}

// 索引可以超过len, 但是不能超越cap
