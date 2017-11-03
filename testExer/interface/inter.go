package main

import (
	"fmt"
)

type int1 int32

type inter1 interface {
	testF1()
}

func (i int1) testF1() {
	fmt.Println("ok")
}

func main() {
	var i1 int1 = 2

	/*
	当一个type满足某一interface时, 可以将其指针赋值给
	interface, 反之则不行, 即: 如果满足interface的是type的
	指针, 赋值给interface的值类型也必须是type的指针, type的
	值是不行的; 可以理解为指针对于interface的适用性更广.
	*/
	var it1 inter1 = &i1
	it1.testF1()
}
