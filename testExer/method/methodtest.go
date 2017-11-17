package main

import (
	"fmt"
)

type vint int32

func (v1 vint) set1() {
	v1++
}

func (v1 *vint) set2() {
	*v1++
}

func main() {
	var vv vint = 10
	fmt.Println(vv)
	vv.set1()
	fmt.Println(vv)
	vv.set2()
	fmt.Println(vv)
	vp := &vv
	vp.set2()
	fmt.Println(vv)
}
