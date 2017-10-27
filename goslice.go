package main

import "fmt"

func modify_slice(a []string) { // 对slice反正操作
	for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Printf("2nd: %q\n", a)
}

func main() {
	s1 := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("1st: %q\n", s1)
	modify_slice(s1)
	fmt.Printf("3th: %q\n", s1)
}
