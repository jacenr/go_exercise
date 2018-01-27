package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
	ENDFOR:
		for {
			select {
			case i := <-c:
				if i == 10 {
					break ENDFOR
				}
				fmt.Println(i)
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("*")
			}
		}
		// var i int
	}()
	for i := 0; i <= 10; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	time.Sleep(2 * time.Second)
}

// select无论有没有default都需要放在for中才能实现轮询
