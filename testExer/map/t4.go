package main

import (
	"fmt"
)

func main() {
	a := make(map[string]map[int]*string) // 把map作为另一个map的value
	b := make(map[int]*string)
	c := "test"
	b[1] = &c
	a["k"] = b
	fmt.Println(a)
}
