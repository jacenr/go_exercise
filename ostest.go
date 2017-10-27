package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	cwDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cwDir)
	fmt.Println(path.Base(os.Args[1]))
	fmt.Println(path.Dir(os.Args[1]))
}
