package main

import (
	"bufio"
	"fmt"
	// "io"
	"os"
)

func main() {
	// r := bufio.NewReader(os.Stdin)

	// var rCount int
	// for {
	// 	txt, err := r.ReadString('\n')
	// 	fmt.Println(txt)
	// 	rCount++
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	// fmt.Println(rCount)

	var sCount int
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	for {
		b := s.Scan()
		fmt.Println(s.Text())
		sCount++
		if !b {
			break
		}
	}
	fmt.Println(sCount)

	// fmt.Scan testing
	// var a string
	// var b int
	// fmt.Scan(&a, &b)
	// fmt.Println(a)
	// fmt.Println(b)
}
