package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := 123
	b := string(a)
	c := strconv.Itoa(a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%b\n", a)
	d := fmt.Sprintf("%b\n", a)
	fmt.Println(d)
	fmt.Fprintf(os.Stdout, "string d: %s\n", d)
	var e bytes.Buffer
	fmt.Fprintf(&e, "%s\n", d)
	fmt.Println(e.String())
}
