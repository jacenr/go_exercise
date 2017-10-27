package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"test1"
)

type testf float64

func (tf testf) Testprint() {
	fmt.Println(tf)
}

func main() {
	var a test1.Testinter
	b := testf(1.26533)
	a = b
	c := test1.Testint(12)
	a.Testprint()
	b.Testprint()
	c.Testprint()
	for _, r := range `\|/` {
		fmt.Printf("%c\n", r)
	}
	// for _, r := range "\|/" {
	// 	fmt.Printf("\r%c\n", r)
	// }
	// fmt.Println(`fdsag\ndsag`)
	// fmt.Println("fdsag\ndsag")
	e := math.Ceil(3.14)
	fmt.Println(e)
	f := math.Ceil(3.0)
	fmt.Println(f)
	m := 4
	// n := 3
	fmt.Println(float64(m) / 3)
	g := "fdsagsagerwa"
	fmt.Println(strings.SplitN(g, "", 3))

	// 字符串排序测试: split + sort + jone(笨方法)
	h := strings.Split(g, "")
	// fmt.Println(strings.Split(g, ""))
	fmt.Println(h)
	sort.Strings(h)
	fmt.Println(h)
	j := strings.Join(h, "")
	fmt.Println(j)
}
