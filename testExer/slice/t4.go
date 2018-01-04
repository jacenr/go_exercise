package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "b", "c", "d"}
	fmt.Println(a)
	fmt.Println(a[1:3])
	fmt.Println(a[1:len(a)])
	a = a[:len(a)-1]
	fmt.Println(a)
}
