package main

import (
	"fmt"
)

type t1 struct {
	i int
}

type t2 struct {
	s string
}

type Ts struct {
	*t1
	*t2
}

type IOne interface {
	f1(x int) int
	f2() string
}

func (a *t1) f1(x int) int {
	return a.i + x
}

func (b *t2) f2() string {
	return fmt.Sprintf("%s", b.s)
}

func main() {
	t1S1 := t1{3}
	// fmt.Println(t1S1.f1(2))
	t2S1 := t2{"Hello Jacky!"}
	// fmt.Println(t2S1.f2())
	tsS1 := Ts{&t1S1, &t2S1}
	// fmt.Println(tsS1)
	var i IOne
	i = tsS1
	fmt.Println(i.f1(3))
	fmt.Println(i.f2())
}
