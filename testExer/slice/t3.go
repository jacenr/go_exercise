package main

import (
	"fmt"
)

func main() {
	// s1 := []int{}
	s1 := make([]int, 0, 1)
	fmt.Printf("%p\n", s1)
	fmt.Println(s1)
	testFunc(s1)
	fmt.Printf("%p\n", s1)
	fmt.Println(s1)
	fmt.Println("===================")
	s2 := []int{1}
	fmt.Printf("%p\n", s2)
	fmt.Println(s2)
	testFunc2(s2)
	fmt.Printf("%p\n", s2)
	fmt.Println(s2)

}

func testFunc(s []int) {
	s = append(s, 1)
}

func testFunc2(s []int) {
	s[0] = 2
}

// slice 数据结构：引用+底层数据
// 当slice作为函数参数时，是吧slice的引用部分的copy传递给了函数
// 在函数内部对slice的引用部分所做的修改，比如append，不会影响到函数外部变量
// 在函数内部对slice的引用所指向的底层数据所做的修改会影响到外部变量，比如对slice元素的修改
// slice本身和指针的作用类似，不同的是在函数内部很少直接修改指针，因为很危险（C语言中可以修改）
