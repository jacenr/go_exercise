package main

import (
	"fmt"
	// "strconv"
	"sync"
)

func main() {
	var x, y int
	var wg sync.WaitGroup
	ch := make(chan string, 2)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x = 1                     // A1
		fmt.Println("y:", y, " ") // A2
		ch <- "done"
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		y = 1                     // B1
		fmt.Println("x:", x, " ") // B2
		ch <- "done"
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	// l := cap(ch)
	// for i := 0; i < l; i++ {
	// 	fmt.Println("ch len: " + strconv.Itoa(len(ch)))
	// 	fmt.Println(<-ch)
	// }
	for s := range ch {
		fmt.Println(s)
	}
}
