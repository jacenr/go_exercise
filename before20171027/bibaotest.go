package main

import (
	"fmt"
)

func main() {
	testfunc()
}

func testfunc() {
	test1 := func() {
		fmt.Println("in bibao 1")
		return
		fmt.Println("in bibao 2")
	}
	fmt.Println("out bibao 1")
	test1()
	// return
	fmt.Println("out bibao 2")
}
