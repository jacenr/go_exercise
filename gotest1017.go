package main

import (
	"fmt"
)

type celu float64

func main() {
	var a celu = 20.0 // ok
	fmt.Printf("%f\n", a)
	fmt.Printf("%T\n", a)
	b := 20.0 // float64
	var c celu
	c = b // error
	fmt.Printf("%f\n", c)
}
