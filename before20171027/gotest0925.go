package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.FreezingC)
	fmt.Println(tempconv.BoilingC)
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
	var k tempconv.Kelvin = 0
	fmt.Println(k)
	fmt.Println(tempconv.KToC(k))
	var a [10]byte
	// for i, _ := range a {
	// for i := range a {
	for i, v := range a {
		fmt.Println(i, v)
	}
}
