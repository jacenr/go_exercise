package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	b, ok := a[0]
	fmt.Println(b)
	fmt.Println(ok)
}

// slice取元素值, 不能通过第二个左值判断
// 须提前判断slice的长度
