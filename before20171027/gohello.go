package main

import (
	"fmt"
)

func main() {

	var pc [16]byte

	for i := range pc {
		//fmt.Println(i)
		//pc[i] = byte(10)
		pc[i] = pc[i/2] + byte(i&1)
	}

	// for i := 0; i <= 7; i++ {
	// 	pc[i] = byte(10)
	// }

	fmt.Println("OK")
	//fmt.Println(byte(2))
	fmt.Println(pc, "end")
	fmt.Println(byte(10 >> 0))
	fmt.Println(byte(10 >> 8))
	fmt.Println(byte(10 >> 16))
	fmt.Println(byte(10 >> 24))
	fmt.Println(byte(10 >> 32))
	fmt.Println(byte(10 >> 40))
	fmt.Println(byte(10 >> 48))
	fmt.Println(byte(10 >> 56))
}
