package main

import (
	"fmt"
)

type Connector interface {
	Connect()
}
type Usb interface {
	Connector
	Name() string
}
type PhoneConnector struct {
	Namec string
}

type NewString string

type StrInter interface {
	output(s string)
}

func StrOut(str StrInter, s string) {
	str.output(s)
}

func (str *NewString) output(s string) {
	*str = *str + NewString(s)
	fmt.Println(*str)
}

func (pc PhoneConnector) Name() string {
	return pc.Namec
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected.")
}

func main() {
	// s := NewString("stringtest1")
	// sp := &s
	s2 := NewString("stringtest2")
	sp := &s2
	// s.output("test")
	// sp.output("test2")
	// StrOut(s, "test1")
	StrOut(sp, "test2")
	var u Usb
	var c Connector
	pc := PhoneConnector{"Phone Connector"}
	u = pc
	// c = pc
	c = u
	fmt.Println(u.Name())
	c.Connect()
}
