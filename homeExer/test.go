package main

import (
	"fmt"
)

func main() {
	s1 := []string{}
	s1 = append(s1, "abc")
	s1 = append(s1, "def")
	s1 = append(s1, "abc")
	fmt.Println(s1)
}
