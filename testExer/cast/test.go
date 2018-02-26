package main

import (
	"fmt"
)

func main() {
	var a int
	a = -8
	fmt.Println(uint(a))
	var b uint
	b = 18446744073709551609
	fmt.Println(int(b))
}
