package main

import (
	"fmt"
)

type Humen struct {
	Name string
	Age  int
}

type Workor struct {
	Humen
	Job string
}

type Thumen Humen

func main() {
	A := Workor{Humen: Humen{"joe", 20}, Job: "work1"}

	H1 := Humen{"joe2", 21}
	B := Workor{H1, "work2"}
	C := Workor{}
	C.Humen = H1
	C.Job = "work3"

	D := Workor{}
	D.Name = "joe3"
	D.Age = 22
	D.Job = "work_sth"

	fmt.Println(A)
	fmt.Println(H1)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)

	H1.printhumen()
	D.printhumen()

	T := Thumen{}
	fmt.Println(T)
	T.printth()
	Humen(T).printhumen()

	E := T.printth // method value
	E()

	F := Thumen.printth // method expression
	F(T)
	// H1.printhumen(3)
}

func (a Humen) printhumen() {
	fmt.Println("1st humen method")
}

func (t Thumen) printth() {
	fmt.Println("1st th method")
}

// func (a Humen) printhumen(i int) {
// 	fmt.Println("2nd method")
// 	fmt.Println(i)
// }

func (w Workor) printhumen() {
	fmt.Println("1st worker method")
}
