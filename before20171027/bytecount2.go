package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	count := 0
	for a > 0 {
		count++
		a &= a - 1
	}
	fmt.Println(count)
}
