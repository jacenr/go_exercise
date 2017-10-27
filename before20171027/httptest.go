package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func main() {
	url := os.Args[1]
	fmt.Println(path.Base(url))
	fmt.Println(url)
	resp, _ := http.Get(url)
	url2 := resp.Request.URL.Path
	fmt.Println(path.Base(url2))
	fmt.Println(url2)
}
