package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	count := 0
	// var i uint
	// for i = 0; i < 64; i++ {
	// 	// a >>= i
	// 	// b := a & 1
	// 	if ((a >> i) & 1) == 1 {
	// 		count++
	// 	}
	// }
	for a > 0 {
		count += (a & 1)
		a = a >> 1
	}
	fmt.Println(count)
}
