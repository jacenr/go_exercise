package main

import (
	"fmt"
)

func f1(arg1 ...interface{}) {
	for _, arg := range arg1 {
		fmt.Println(arg)
	}
}

func main() {
	f1("a", 1)
}
