package main

import (
	// "bytes"
	// "crypto/sha256"
	"fmt"
	"strconv"
)

// func bitcount(){}

func main() {
	// sx := sha256.Sum256([]byte{'x'})
	// sx16 := fmt.Sprintf("%x", sx)
	// fmt.Printf("%b\n", sx16)
	sx10, err := strconv.ParseInt("a", 16, 64)
	fmt.Println(sx10)
	fmt.Println(err)

	// var buf bytes.Buffer
	// sx := sha256.Sum256([]byte{'x'})
	// fmt.Println(sx)
	// fmt.Println(len(sx))
	// fmt.Printf("%x\n", sx)
	// a := fmt.Sprintf("%b", sx)
	// // fmt.Println(a)
	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("a type: %T\n", a)
	// sx2 := sx[:]
	// // sx3 := []byte{1, 2, 3}
	// fmt.Printf("%v\n%T\n", sx2, sx2)
	// n, err := buf.Write(sx2)
	// fmt.Printf("%v\n", buf.String())
	// fmt.Println(n, err)
	// // fmt.Printf("%s\n", sx)
}
