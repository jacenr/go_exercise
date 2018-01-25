package main

import (
	"fmt"
)

func main() {
	var a []string
	fmt.Println(a == nil)
	fmt.Println(len(a))
	b := make([]string, 1)
	fmt.Println(b == nil)
	fmt.Println(len(b))
	c := []string{}
	fmt.Println(c == nil)
	fmt.Println(len(c))
}

// slice为nil时, 长度必为0, slice不为nil时, 长度可以是0, 也可以不为0
