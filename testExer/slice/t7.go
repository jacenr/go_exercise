package main

import (
	"fmt"
)

func main() {
	a := make([]chan string, 0, 3)
	for i := 0; i < 3; i++ {
		b := make(chan string)
		a = append(a, b)
	}
	fmt.Println(a)
}

// channel可以作为slice的元素
