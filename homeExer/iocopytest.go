package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// input :=
	fmt.Println("ok")
	io.Copy(os.Stdout, os.Stdin)
}
