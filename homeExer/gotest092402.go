package main

import (
	"fmt"
	// "log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homehandle)
	http.ListenAndServe("localhost:8080", nil)
}

func homehandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "url path: %s", r.URL.Path)
}
