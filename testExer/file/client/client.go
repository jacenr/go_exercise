package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	// "io"
	"net"
	"os"
)

type mg struct {
	List []string
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	fmt.Println(err)
	cnrd := bufio.NewReader(conn)
	dec := gob.NewDecoder(cnrd)
	var m mg
	err = dec.Decode(&m)
	fmt.Println(err)
	f, err := os.Create("tmp2")
	fmt.Println(err)
	defer f.Close()
	for _, v := range m.List {
		f.WriteString(v)
	}
}
