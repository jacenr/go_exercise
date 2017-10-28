package main

import (
	"fmt"
	"strings"
)

func MapString(s string) map[string]int {
	ss := strings.Split(s, " ")
	sm := make(map[string]int)
	for _, v := range ss {
		sm[v] += 1
	}
	return sm
}

func main() {
	s := "fds fsfgre grehtth tr jy jyyr e het fds tr"
	fmt.Println(MapString(s))
}
