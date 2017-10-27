package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	sha := flag.String("type", "sha256", "Please give the sha type.")
	flag.Parse()
	var a []byte
	switch *sha {
	case "sha512":
		tmp := sha512.Sum512([]byte{'x'})
		a = tmp[:]
	case "sha384":
		tmp := sha512.Sum384([]byte{'x'})
		a = tmp[:]
	default:
		tmp := sha256.Sum256([]byte{'x'})
		a = tmp[:]
	}
	fmt.Println(a)
}
