package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	urls := os.Args[1:]
	ch := make(chan string)
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		// res, err := http.Get(url)
		go handleurl(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}

}

func handleurl(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	end := time.Since(start).Seconds()
	if err != nil {
		ch <- fmt.Sprint(err)
		// os.Exit(1)
		return
	}
	n, err := io.Copy(ioutil.Discard, res.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("%d\t%f", n, end)
	res.Body.Close()
}
