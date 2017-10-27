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

func main() {
	s1 := []string{}
	fmt.Printf("函数外部append之前的值: %s\n", s1)
	test1(s1)
	fmt.Printf("函数外部append之后的值: %s\n", s1)

}
