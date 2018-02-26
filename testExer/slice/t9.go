package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := a[0:len(a)] // slice的二次切片不包含最后一个索引
	fmt.Println(b)
}
