package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{}
	m2 := make(map[int]string)
	m1[1] = "abc"
	m2[1] = "abc"
	fmt.Println(m1 == nil)
	fmt.Println(m2 == nil)
	fmt.Println(m1)
	fmt.Println(m2)
	var m3 map[int]string
	fmt.Println(m3 == nil)
	m3 = map[int]string{}
	fmt.Println(m3 == nil)
}
