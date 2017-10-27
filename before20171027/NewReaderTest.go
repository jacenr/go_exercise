package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// NewReader
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, _ := in.ReadRune()
		// fmt.Println(r)
		fmt.Printf("char: %c\t%U\n", r, r)
	}

	// NewScanner
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	fmt.Printf("%v\n", input.Text())
	// }
}
