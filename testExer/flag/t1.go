package main

import (
	"flag"
	"fmt"
)

func main() {
	a := flag.String("a", "", "A value.")
	flag.Parse()
	fmt.Println(*a)
	fmt.Println(*a == "")
}
