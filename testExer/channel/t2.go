package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	v, ok := <-ch
	fmt.Printf("%v %d\n", ok, v)
	v, ok = <-ch
	fmt.Printf("%v %d\n", ok, v)
	v, ok = <-ch
	fmt.Printf("%v %d\n", ok, v)
	v, ok = <-ch
	fmt.Printf("%v %d\n", ok, v)
}

// channel关闭后, 还是可以读出已经写入其中的数据的;
