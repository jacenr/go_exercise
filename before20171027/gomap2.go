package main

import (
	"fmt"
)

func main() {
	mapA := map[int]string{
		1: "a",
		2: "b",
	}
	fmt.Println(mapA)
	modifyMap(mapA)
	fmt.Println(mapA)
}

func modifyMap(m map[int]string) {
	m = nil
}
