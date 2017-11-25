package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	lst, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := lst.Accept()
		defer conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
		rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			// rw.WriteString("test" + strconv.Itoa(i) + "\n")
			rw.Write([]byte("test" + strconv.Itoa(i) + "\n"))
			rw.Flush()
			time.Sleep(1 * time.Second)
		}
		// rw.Flush()
	}
}