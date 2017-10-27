package main

import (
	"fmt"
)

func test1(s []string) {
	fmt.Printf("函数内部初始长度: %d\n", len(s))
	fmt.Printf("函数内部初始值: %s\n", s)
	s = append(s, "a")
	fmt.Printf("函数内部修改后长度: %d\n", len(s))
	fmt.Printf("函数内部修改后值: %s\n", s)
}

func test2(s []string) {
	fmt.Printf("函数内部swap之前的初始值: %s\n", s)
	n := len(s) - 1
	for i, j := 0, n; i <= j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Printf("函数内部swap之后append之前的值: %s\n", s)
	s = append(s, "d")
	fmt.Printf("函数内部swap及append之后的值: %s\n", s)
}

func test3(s []string) {
	fmt.Printf("函数内部原地修改之前的值: %s\n", s)
	for i, v := range s {
		s[i] = v + v
	}
	fmt.Printf("函数内部原地修改后的值: %s\n", s)
	s = append(s, "d")
	fmt.Printf("函数内部原地修改及append之后的值: %s\n", s)
}

func main() {
	s1 := []string{}
	fmt.Printf("函数外部append之前的值: %s\n", s1)
	test1(s1)
	fmt.Printf("函数外部append之后的值: %s\n", s1)
	fmt.Println("===================")
	s2 := []string{"a", "b", "c"}
	fmt.Printf("函数外部swap之前的值: %s\n", s2)
	test2(s2)
	fmt.Printf("函数外部swap及append之后的值: %s\n", s2)
	fmt.Println("===================")
	s3 := make([]string, 0, 4)
	fmt.Printf("函数外部初始带容量的值: %s\n", s3)
	test1(s3)
	fmt.Printf("函数外部初始带容量的append之后的值: %s\n", s3)
	fmt.Println("===================")
	s4 := []string{"a", "b", "c"}
	fmt.Printf("函数外部原地修改之前的值: %s\n", s4)
	test3(s4)
	fmt.Printf("函数外部原地修改之后的值: %s\n", s4)
}
