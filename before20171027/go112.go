package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form Error: %v\n", err)
		return
	}
	for k, v := range r.Form {
		if k == "cycle" {
			rd, _ := strconv.Atoi(v[0])
			area := 3.14 * float64(rd) * float64(rd)
			fmt.Fprintf(w, "The Area is: %f", area)
			return
		} else {
			// fmt.Fprintf(w, "Error: No cycle is found.")
			continue
		}
		// fmt.Fprintf(w, "Form[%s]\t%s\n", k, v)
	}
	fmt.Fprintf(w, "Error: No cycle is found!")
}
