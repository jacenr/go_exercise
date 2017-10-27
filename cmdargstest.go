package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// for _, a := range os.Args[1:] {
	// 	fmt.Println(a)
	// }
	var aint = flag.Int("i", 2, "Please give a int number.")
	var addr = flag.String("NewYork", "localhost:8080", "Please give a ip addr with a port.")
	flag.Parse()
	fmt.Println(*aint)
	fmt.Println(*addr)
	for i, a := range os.Args[3:] {
		fmt.Println(i, a)
		m := strings.Split(a, "=")
		type ninter interface{}
		var n ninter
		n = m
		mtype, err := n.([]string)
		fmt.Println(mtype, err)
	}
}
