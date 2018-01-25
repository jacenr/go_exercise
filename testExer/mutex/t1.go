package main

import (
	"fmt"
	"sync"
	"time"
)

var v1 int32 = 10

type st struct {
	m  sync.Mutex
	v2 int32
}

var s1 = st{v2: 10}

func main() {
	var wg sync.WaitGroup

	// 测试1
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		go func() {
			s1.m.Lock() // struct内嵌的mutex可以实现对自身数据的加锁
			defer s1.m.Unlock()
			t1 := s1.v2
			t1++
			time.Sleep(2 * time.Second) // sleep是为了能够看到数据竞争的结果
			s1.v2 = t1
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(s1.v2)

	// 测试2
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		go func() {
			s1.m.Lock()
			defer s1.m.Unlock() // struct内嵌的mutex也可以实现对其他数据的加锁
			t1 := v1
			time.Sleep(1 * time.Second) // sleep是为了能够看到数据竞争的结果
			t1++
			v1 = t1
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(v1)
}

// 测试结果: 结构中内嵌sync.Mutex仅仅是作为struct的一个成员而已, 其即可以对
// struct其他成员加锁, 也可以对非struct成员的数据加锁

// mutex是个struct, 它具有Lock()和Unlock()方法, 同一时刻只能由一个goroutine
// 对同一个mutex施加lock, 其他goroutine就会一直阻塞至mutex nulock, 以达到对
// mutex之后的代码加锁的目的.
