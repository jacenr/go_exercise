package main

import (
	"fmt"
)

var n int

func foo() (int, error) {
	return 5, nil
}

func bar() {
	fmt.Println("bar n:", n)
}

func main() {
	n, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println("main n:", n)
}
