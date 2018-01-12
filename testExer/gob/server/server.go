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

type t2 struct {
	I int
}

func main() {
	lst, _ := net.Listen("tcp", "127.0.0.1:8999")
	for {
		conn, _ := lst.Accept()
		cnwt := bufio.NewWriter(conn)
		enc := gob.NewEncoder(cnwt)
		v1 := t1{"test1"}
		err := enc.Encode(v1)
		fmt.Println(err)
		cnwt.Flush()
		fmt.Println(v1)
		v2 := t1{"test2"}
		enc.Encode(v2)
		cnwt.Flush()
		fmt.Println(v2)
		v3 := t2{1}
		enc.Encode(v3)
		cnwt.Flush()
		fmt.Println(v3)
		conn.Close()
		// break
	}
}
