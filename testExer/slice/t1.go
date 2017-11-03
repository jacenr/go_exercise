package main

import (
	"fmt"
)

func main() {
	// 对二次切片元素的修改会影响到一次切片
	// 慎用二次切片
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[2:5]
	s3 := []int{3, 4, 5}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	s2[2] = 7
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
