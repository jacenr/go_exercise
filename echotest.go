package main

import (
	"fmt"
	"os"
	// "strings"
	"strconv"
)

// 输出 程序名+参数
// func main() {
// 	fmt.Println(strings.Join(os.Args[:], " "))
// }

// 输出 每个参数索引+参数 每个一行
func main() {
	for ind, arg := range os.Args[:] {
		s := strconv.Itoa(ind) + " " + arg
		//fmt.Println(ind)
		fmt.Println(s)
	}
}
