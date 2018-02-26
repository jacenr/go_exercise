package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("file1.txt")
	r := bufio.NewReader(f)
	p := make([]byte, 8)
	for {
		n, err := r.Read(p)
		fmt.Printf("%d %v\n", n, err)
		for _, c := range p {
			fmt.Printf("%c", c)
		}
		fmt.Println()
		if err == io.EOF {
			fmt.Println(n)
			break
		}
	}
}
