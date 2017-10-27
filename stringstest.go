package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	urlstrs := os.Args[1:]
	for _, urlstr := range urlstrs {
		strSlc := strings.Split(urlstr, ".")
		l := len(strSlc)
		DomainStr := strings.Join(strSlc[l-2:], ".")
		fmt.Println(DomainStr)
	}
}
