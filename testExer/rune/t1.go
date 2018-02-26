package main

import (
	// "bufio"
	// "bytes"
	"fmt"
	"io"
	"os"
	// "unicode/utf8"
)

func main() {
	srcF, _ := os.Open("test1.txt")
	dstF, _ := os.Create("test2.txt")

	dstF2, _ := os.Create("test3.txt")
	// ** io.Copy **
	n, _ := io.Copy(dstF, srcF) // io.Copy直接copy中文字符没有问题。
	fmt.Println(n)
	sn, _ := srcF.Seek(0, 0)
	fmt.Println(sn)
	n2, _ := io.Copy(dstF2, srcF) // io.Copy直接copy中文字符没有问题。
	fmt.Println(n2)

	// buf := bufio.NewReader(srcF)
	// var line rune
	// var err error

	// for {
	// 	line, _, err = buf.ReadRune()
	// 	// fmt.Println(line)
	// 	fmt.Printf("%q\n", line)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

	srcF.Close()
	dstF.Close()
	dstF2.Close()
}
