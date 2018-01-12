package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
)

type t1 struct {
	S string
}

func main() {
	var vt1 t1
	conn, _ := net.Dial("tcp", "127.0.0.1:8999")
	cnrd := bufio.NewReader(conn)
	dec := gob.NewDecoder(cnrd)
	for i := 0; i < 3; i++ {
		err := dec.Decode(&vt1)
		fmt.Println(err)
		fmt.Println(vt1)
	}
	conn.Close()
}
