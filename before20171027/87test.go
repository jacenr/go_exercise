package main

import (
	"fmt"
	"os"
	"time"
)

// func main() {
// 	fmt.Println("Commencing countdown.")
// 	tick := time.Tick(1 * time.Second)
// 	fmt.Println(len(tick))
// 	fmt.Println(cap(tick))
// 	for countdown := 10; countdown > 0; countdown-- {
// 		fmt.Println(countdown)
// 		// fmt.Println(<-tick)
// 		<-tick
// 	}
// 	// launch()
// }

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	// ...create abort channel...
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// fmt.Println("time")
		// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	// launch()
}
