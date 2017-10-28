package main

import (
	"fmt"
	"os"
	// "strconv"
	// "strings"
	"time"
)

func main() {
	start := time.Now().Unix()
	s, sep := "", ""
	// s = strings.Join(os.Args, sep)
	// time.Sleep(3 * time.Second)
	for _, v := range os.Args {
		s += sep + v
		sep = " "
	}
	end := time.Now().Unix()
	fmt.Println(s)
	// time_p = end - start
	fmt.Println(start)
	fmt.Println(end)
	// for i, v := range os.Args {
	// 	fmt.Println(strconv.Itoa(i) + "," + v)
	// }
}
