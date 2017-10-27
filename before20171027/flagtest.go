package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	bo := flag.Bool("b", false, "give the bool value")
	fi := flag.String("f", "", "give the file name")
	te := flag.String("t", "", "give the test name")
	flag.Parse()
	fmt.Println("the all args:", flag.Arg(0))
	fmt.Println("the all args:", flag.Arg(1))
	fmt.Println("the all args:", flag.Args())
	fmt.Println("the bool value:", *bo)
	fmt.Println("the file name:", *fi)
	fmt.Println("the file name:", *te)
	count := len(os.Args)
	for i := 0; i < count; i++ {
		fmt.Println(i, ":", os.Args[i])
	}
}
