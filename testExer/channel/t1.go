package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 3)

	// **测试同时读取。**
	// go func() {
	// 	time.Sleep(2 * time.Second) // 等待所有goroutine准备就绪。
	// 	ch <- 1
	// 	ch <- 2
	// 	ch <- 3
	// }()
	// for i := 0; i <= 2; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		fmt.Println(<-ch)
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// **测试同时写入**
	ch_wait := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second) // 等待所有goroutine准备就绪。
		close(ch_wait)              // 广播
	}()
	for i := 0; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-ch_wait
			ch <- i
		}(i)
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("ok")
}

// 测试表明1：只要写入没问题，且channel缓存个数>=goroutine的数目，多个goroutine同时读取
// 同一个channel是没有问题的，无须加锁。

// 测试表明2：只要channel缓存个数>=goroutine数目，即写入没问题，多个goroutine可以同时对
// 同一个channel写入数据是没有问题的，无须加锁。

// 虽然多个goroutine同时读写同一个channel在操作上是没有问题，但是据说对性能有影响。
