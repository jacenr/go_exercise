package main

import (
	"fmt"
)

type as struct {
	s1 string
}

type bs struct {
	*as
	s2 string
}

func (a *as) f1() {
	fmt.Println(a.s1)
}

func main() {
	a := &as{"a1"}
	b := bs{a, "b1"}
	a.f1()
	b.f1()
}
