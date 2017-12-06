package main

import (
	"fmt"
	"os"
	// "time"
)

func main() {
	file := os.Args[1]
	fi, _ := os.Lstat(file)
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.Mode())
	fmt.Println(fi.IsDir())
	fmt.Println(fi.Sys())
}
