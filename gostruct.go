package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func modify_struct1(a *Employee) {
	a.ID = 654321
	a.Name = "my 2nd name"
}

func main() {
	var dilbert Employee

	dilbert.ID = 123456
	dilbert.Name = "my 1st name"

	dilbertP := &dilbert

	fmt.Println(dilbert.ID)
	fmt.Println(dilbert.Name)
	modify_struct1(dilbertP)
	fmt.Println(dilbert.ID)
	fmt.Println(dilbert.Name)
	// dilbert.Salary = 1000
	// fmt.Println(dilbert.Salary)
	// dbt := &dilbert
	// fmt.Println((*dbt).Salary)
	// fmt.Println(dbt.Salary)
	// dbt.ID = 9
	// fmt.Println(dilbert.ID)
	// fmt.Println(dbt.ID)

}
