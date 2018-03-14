package main

import (
	"fmt"
	"math/rand"
	"time"
)

type st struct {
	n int
	p *st
}

func main() {
	now := time.Now()
	rand.Seed(now.UnixNano())
	t := new(st)
	t0 := t
	fmt.Println("===========first===========")
	for i := 0; i < 10; i++ {
		t1 := st{rand.Int(), nil}
		fmt.Println(t1.n)
		t0.p = &t1
		t0 = t0.p
	}
	fmt.Println("===========again===========") // 对比
	t0 = t
	for {
		if t0.p == nil {
			break
		}
		t0 = t0.p
		fmt.Println(t0.n)
	}
	fmt.Println("===========reverse===========")
	t3 := new(st)
	t0 = t.p
	t2 := t0.p
	t0.p = nil
	for {
		t3 = t2.p
		t2.p = t0
		t0 = t2
		t2 = t3
		if t3 == nil {
			break
		}
	}
	for {
		fmt.Println(t0.n)
		if t0.p == nil {
			break
		}
		t0 = t0.p
	}
}
