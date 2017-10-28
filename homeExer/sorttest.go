package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func main() {
	testSlice := []string{"gras", "re", "ktfhds"}
	sort.Sort(StringSlice(testSlice))
	fmt.Println(testSlice)
}

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
