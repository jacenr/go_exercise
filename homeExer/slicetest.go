package main

import (
	"fmt"
)

func main() {
	a1 := [...]int{2, 5, 58, 6, 9, 7, 5, 69, 9, 8, 5, 9}
	s1 := a1[:]
	s2 := make([]int, 3, 6)
	s2 = a1[2:6]
	s3 := a1[0:3]
	fmt.Println(cap(s1), len(a1))
	fmt.Printf("%p\t%v\n", s1, s1)
	fmt.Printf("%p\t%v\n", s2, s2)
	fmt.Printf("%p\t%v\n", s3, s3)
}
