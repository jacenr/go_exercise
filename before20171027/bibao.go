package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	fmt.Printf("start a=%d, b=%d\n", a, b)
	func(a int) { a = 11 }(a)
	// panic("test")
	func() { b = 22 }()
	fmt.Printf("end a=%d, b=%d\n", a, b)
}
