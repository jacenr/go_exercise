package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	modifySlice(a)
	fmt.Println(a)
}

func modifySlice(data []int) {
	data = nil
	fmt.Println(data)
}
