package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Symlink("file1", "slink1")
	if err != nil {
		fmt.Println(err)
	}
	err := os.Symlink("file2", "slink1")
	if err != nil {
		fmt.Println(err)
	}
}
