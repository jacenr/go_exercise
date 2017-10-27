package main

import (
	"fmt"
)

func main() {
	MapA := make(map[int]string)
	MapA[1] = "a"
	MapA[2] = "b"
	fmt.Println(MapA)
	modifyMap(MapA)
	fmt.Println(MapA)
}

func modifyMap(m map[int]string) {
	m[100] = "xyz"
}
