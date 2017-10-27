package main

import (
	"fmt"
	"strings"
	"time"
)

func add1(r rune) rune { return r + 1 }

func main() {
	// var a rune = "HAL-9000"
	// fmt.Printf("%x", a)
	// a := add1("HAL-9000")
	defer testFunc("test a")()
	time.Sleep(10 * time.Second)
	fmt.Println(strings.Map(add1, "HAL-9000"))
	// fmt.Println(a)
	// for _, charTest := range "HAL-9000" {
	// 	fmt.Printf("%q\n", add1(charTest))
	// }
}

func testFunc(s string) func() {
	start := time.Now()
	fmt.Printf("enter %s: %v\n", s, start)
	return func() {
		fmt.Printf("exit %s: %v, %v\n", s, time.Now(), time.Since(start))
	}
}

func(n *html.Node) {
    if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
        if title != "" {
            panic(bailout{}) // multiple titleelements
        }
        title = n.FirstChild.Data
    }
}