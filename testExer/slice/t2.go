package main

import (
	"fmt"
)

type ip []byte

func main() {
	var i1 ip = ip("test")
	fmt.Println(i1)
	fmt.Printf("%T\n", i1)
	b1 := []byte(i1)
	fmt.Println(b1)
	fmt.Printf("%T\n", b1)
}
