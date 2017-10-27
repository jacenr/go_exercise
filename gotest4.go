package main

import (
	"bytes"
	"fmt"
)

type testS struct{ int }

type testS2 struct{ Interface }

type Interface interface{}

func main() {
	var a, b *bytes.Buffer
	b = new(bytes.Buffer)
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Printf("%T\t%T\n", a, b)
	c := "a"
	fmt.Println(c[0])
	fmt.Printf("%c\n", c[0])
	fmt.Printf("%s\n", c)
	fmt.Println(byte(c[0]))
	d := testS{3}
	fmt.Printf("%v\n", d)
	e := testSI(6)
	fmt.Printf("%T\t%v\n", e, e)
	var f Interface
	g := testSIn(f)
	fmt.Printf("%T\n", g)
	h := 1.335246
	fmt.Println(int(h))
}

func testSI(i int) testS { // 返回Interface类型竟然也可以
	return testS{i}
}

func testSIn(i Interface) Interface {
	return testS2{i}
}
