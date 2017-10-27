package main

import (
	"fmt"
)

func main() {
	a := Max(2, 56, 5, 9, 6, 35, 9, 6, 8)
	fmt.Println(a)
}

func Max(vars ...int) int {
	if len(vars) == 0 {
		fmt.Println("Please give more than 0 args.")
		return 0
	}
	var maxInt int = vars[0]
	for _, i := range vars {
		if i > maxInt {
			maxInt = i
		}
	}
	return maxInt
}
