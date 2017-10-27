package main

import (
	"fmt"
)

func main() {
	for _, v := range test(nil) {
		fmt.Println(v)
	}
	// var sll []string
	// fmt.Println(sll == nil)
	// sll = append(sll, "test")
	// fmt.Println(sll)
	// fmt.Println(sll == nil)
}

// 即使slice是nil, 也可以向其append值
// 这个函数可以接受nil的参数, 因为其会返回一个slice, 所以外层函数还是可以访问到return的值的
func test(sl []string) []string {
	// var sl []string
	// fmt.Println(sl == nil)
	sl = append(sl, "a")
	sl = append(sl, "b")
	return sl
}
