package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	// "unicode/utf8"
)

func main() {
	srcF, _ := os.Open("test1.txt")
	// dstF, _ := os.Create("test2.txt")

	// ** io.Copy **
	// n, _ := io.Copy(dstF, srcF) // io.Copy直接copy中文字符没有问题。
	// fmt.Println(n)

	buf := bufio.NewReader(srcF)
	var line []byte
	var err error

	for {
		line, err = buf.ReadBytes('\n')
		// fmt.Println(line)
		fmt.Printf("%v\n", bytes.Runes(line))
		if err == io.EOF {
			break
		}
	}

	srcF.Close()
	// dstF.Close()
}
