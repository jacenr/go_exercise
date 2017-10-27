// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "unicode"
	// "unicode/utf8"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	// var utflen []int               // count of lengths of UTF-8 encodings
	// utflen := make([]int, 300)
	// invalid := 0                    // count of invalid UTF-8 characters
	f, _ := os.Open("golang.txt")
	in := bufio.NewScanner(f)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		if in.Err() == io.EOF {
			break
		}
		if in.Err() != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", in.Err())
			os.Exit(1)
		}
		// if r == unicode.ReplacementChar && n == 1 {
		// 	invalid++
		// 	continue
		// }
		word := in.Text()
		counts[word]++
		// counts[r]++
		// utflen[len(word)]++
	}
	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%v\t%d\n", c, n)
	}
	// fmt.Print("\nlen\tcount\n")
	// for i, n := range utflen {
	// 	if i > 0 {
	// 		fmt.Printf("%d\t%d\n", i, n)
	// 	}
	// }
	// if invalid > 0 {
	// 	fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	// }
}
