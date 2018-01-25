package main

import (
	"fmt"
)

var m map[string]string

func main() {
	m = make(map[string]string)
	m["1"] = "a"
	m["2"] = "b"
	m["3"] = "c"
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Println("============")
	m = make(map[string]string)
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}

// 使用新的变量重新赋值可以初始化原来的变量
