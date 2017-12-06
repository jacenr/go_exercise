package main

import (
	"fmt"
	"time"
)

func main() {
	timeForm := "20060102150405"
	timeS, _ := time.Parse(timeForm, "20171206114130")
	fmt.Println(timeS)
	// timeS:=
}
