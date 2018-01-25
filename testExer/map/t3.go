package main

import (
	"fmt"
)

type st struct {
	a string
	b int
}

func main() {
	m := make(map[string]*st)
	m1 := &st{"a", 1}
	m2 := &st{"b", 2}
	m["m1"] = m1
	m["m2"] = m2
	for k, v := range m {
		fmt.Printf("%s: %v\n", k, v)
	}
}

// 指针可以作为map的value
