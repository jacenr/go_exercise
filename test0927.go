package main

import (
	"fmt"
)

func main() {
	var a int
	var b string
	n, err := fmt.Scanf("%d%s", &a, &b)
	fmt.Println(n, err)
	fmt.Println(a)
	fmt.Println(b)
	c := fmt.Sprintf("%T", a)
	fmt.Println(c)
}
