package main

import (
	"fmt"
	"strings"
)

func main() {
	a := ""
	b := strings.Split(a, ",,")
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(b[0])
	fmt.Println(b[0] == "")
	// fmt.Println(b[1])
	fmt.Println(b[0] == " ")
}
