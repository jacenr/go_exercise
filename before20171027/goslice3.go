package main

import "fmt"

func modifySlice(p *[]string) {
	(*p)[0] = "m"
	fmt.Println(*p)
}

func modifySlice2(p *[]string) {
	*p = nil
	fmt.Println(*p)
}

func main() {
	sliceA := []string{"a", "b", "c", "d"}
	fmt.Println(sliceA)
	p := &sliceA
	modifySlice(p)
	fmt.Println(sliceA)
	modifySlice2(p)
	fmt.Println(sliceA)
}
