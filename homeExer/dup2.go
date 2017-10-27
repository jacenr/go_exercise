// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for fileName, count := range counts {
		for line, n := range count {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", fileName, n, line)
			}
		}

	}
}
func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	count := make(map[string]int)
	counts[f.Name()] = count
	for input.Scan() {
		count[input.Text()]++
	} // NOTE: ignoring potential errors from input.Err()
}
