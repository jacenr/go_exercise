package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

type Manager struct {
	User
	Title string
}

func (m Manager) Output() string {
	return m.Name + "," + strconv.Itoa(m.Age) + "," + m.Title
}

func main() {
	u1 := User{"jerry", 20}
	m1 := Manager{u1, "CEO"}
	fmt.Println("1. " + m1.Output())

	t := reflect.TypeOf(m1)
	v := reflect.ValueOf(m1)

	fmt.Println("==========t & v=========")
	fmt.Printf("t: %v\n", t)
	fmt.Printf("v: %v\n", v)
	fmt.Println("========================")
	tl := t.NumField()
	// vl := v.NumField()
	// allField := []string{}
	fmt.Println(tl)
	for i := 0; i < tl; i++ {
		// func() {
		// 	defer func() { recover() }()
		// 	fmt.Println(t.FieldByIndex([]int{i, 0}))
		// }()
		fmt.Println(t.FieldByIndex([]int{1}))
	}
	fmt.Println(t.Field(0).Type.Field(0))

}
