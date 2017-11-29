package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println(io.EOF) // EOF是error类型
	// c := io.EOF
	b := []byte{'\000', '\060', '0', ' ', '\x30'} // \000 or \x00表示字符NULL, \060 or \x30表示字符0，32是空格的十进制数
	fmt.Println(b)
	// fmt.Println(b[1])
	for _, i := range b {
		// fmt.Println(i)
		fmt.Printf("d: %d\n", i)
		fmt.Printf("c: %c\n", i)
		fmt.Println("---------")
	}
}
