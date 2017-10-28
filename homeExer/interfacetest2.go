package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

type Manager struct {
	User
	Title string
}

func (u User) Output() {
	fmt.Println("Name: ", u.Name)
	fmt.Println("Age: ", u.Age)
}

func main() {
	u1 := User{"jerry", 20}
	u1.Output()

	t := reflect.TypeOf(u1)
	v := reflect.ValueOf(u1)
	fmt.Println(" ")
	fmt.Println("t: ", t)
	fmt.Printf("t: %v\n", t)
	fmt.Printf("%v\n", t.NumField())
	fmt.Printf("%v\n", t.Field(0))
	fmt.Printf("%v\n", t.Field(1))
	fmt.Printf("%v\n", t.NumMethod())
	s1 := []int{1}
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", t.FieldByIndex(s1))
	fmt.Println(" ")
	fmt.Println("v: ", v)
	fmt.Printf("v: %v\n", v)
	fmt.Println(v.NumField())
	fmt.Println(v.Field(0))
	fmt.Println(v.Field(1))
	fmt.Println(v.NumMethod())
	fmt.Println(v.Kind())
	fmt.Println(" ")
	fmt.Println(v.Type() == t)
}
