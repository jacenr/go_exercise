package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// if input.Text() == "\n" {
		// 	fmt.Println("OK")
		// 	break
		// }
		if input.Text() != "\n" {
			counts[input.Text()]++
		}
		//fmt.Printf("%s\n", input.Text())
	}

	for line, line_no := range counts {
		fmt.Printf("%d,%s\n", line_no, line)
	}
}
