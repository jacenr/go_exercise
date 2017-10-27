package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 5, 6}
	b := []int{5, 9, 363, 9, 6}
	a = append(a, b...)
	fmt.Println(a)
}
