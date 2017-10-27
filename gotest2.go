package main

import (
	"fmt"
	// "unicode/utf8"
)

func main() {
	a := "\x38"
	b := "\070"
	c := "abcd\x45\u4e26世界\u4e16"
	d := `abcd\x45
  test    a\n` + "`" + `\fdsag\n\tfdsaf`
	e := `abcd\x45test`
	f := "\xe4\xb8\x96\xe7\x95\x8c\u4e07\u30d7ロ\u4eac"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%s\n", d)
	fmt.Printf("%s\n", e)
	fmt.Println(f)
	m := []int{1, 2, 3}
	// n := string(m)
	fmt.Printf("%v\n", m)
	// fmt.Println(n)
}
