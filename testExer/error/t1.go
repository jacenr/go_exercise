package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := errors.New("test text")
	// var e2 error
	fmt.Println(e1.Error())
	// fmt.Println(e2.Error())
}
