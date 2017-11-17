package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func f1(ch1 chan int, ch2 chan string) {
	defer wg1.Done()
	fmt.Println(<-ch1)
	ch2 <- "Done"
}

func f2(ch1 chan int) {
	time.Sleep(1 * time.Second)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 3)
	// go f1(ch1, ch2)
	// go f2(ch1)
	// <-ch2

	// i为goroutine的数目
	for i := 1; i <= 3; i++ {
		wg1.Add(1)
		go f1(ch1, ch2)
	}
	f2(ch1)
	go func() {
		wg1.Wait()
		close(ch2)
	}()
	for i := range ch2 {
		fmt.Println(i)
	}
}

// 测试表明多个goroutine同时读取同一个channel, 不需要加锁, 没有出现死锁的情况
// 当同时读取channel的goroutine数目大于channel的缓存时会出现死锁, 反之不会出现死锁
