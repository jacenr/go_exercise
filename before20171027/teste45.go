package main

import (
	"fmt"
)

func main() {
	strs := []string{"abc", "def", "def", "abcd", "mnj", "mnj", "jkl", "weio"}
	// strs := []string{"abc", "abcd"}
	fmt.Println(strs)
	fmt.Println(removedup(strs))
}

// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

func removedup(strs []string) []string {
start:
	n := len(strs)
	if n == 1 {
		return strs
	}
	for i := 0; i < n-1; i++ {
		if strs[i] == strs[i+1] {
			copy(strs[i:], strs[i+1:])
			strs = strs[:n-1]
			goto start
		}
	}
	return strs
}
