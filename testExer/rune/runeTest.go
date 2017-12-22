package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "中文"
	fmt.Printf("%T\n", a)
	fmt.Println(len(a))

	// 计算中文string的长度需要使用下面的函数
	fmt.Println(utf8.RuneCountInString(a))

	// 对于中文string, 可以直接使用range遍历字符串
	for _, s := range a {
		fmt.Printf("%T\t%c\n", s, s)
	}

	fmt.Println("==================")
	r1 := rune(65)
	fmt.Println(r1)
	fmt.Println(string(r1))
	r2 := []rune{65, 66, 67}
	fmt.Println(string(r2))
	// fmt.Println('A')
	fmt.Println("===================")
	fmt.Printf("%X\t%X\n", ' ', ',')
	fmt.Println("===================")
	str1 := "testtest"
	fmt.Println(len([]rune(str1)))
	fmt.Println(len([]byte(str1)))
	str2 := "你好世界"
	fmt.Println(len([]rune(str2)))
	fmt.Println(len([]byte(str2)))
	// rune适用范围貌似比byte更广。
}
