package main

import (
	"bufio"
	// "bytes"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for {
		b := s.Scan()
		if !b {
			break
		}
		// fmt.Printf("%v\n", b)
		// fmt.Printf("%v\n", s.Err())
		fmt.Println(s.Text())
	}
	fmt.Printf("%v", s.Err())
	// s.Scan()
	// t := s.Text()
	// b := []byte(t)
	// _, tk, _ := bufio.ScanWords(b, true)
	// buf := bytes.NewBuffer(tk)
	// fmt.Println(buf.String())
}
