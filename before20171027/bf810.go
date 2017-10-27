package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	// "tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	fwalk(t, ch)
	close(ch)
}

func fwalk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		fwalk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		fwalk(t.Right, ch)
	}

}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
		// fmt.Println(i)
	}
	// for i := range ch2 {
	// 	fmt.Println(i)
	// }
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t1))
}
