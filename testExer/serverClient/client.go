package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	fmt.Printf("\\n: %d\n", '\n')
	for {
		// str, err := rw.ReadString('\n')
		str, err := rw.ReadBytes('\n')
		if err != nil {
			switch err {
			case io.EOF:
				log.Println("READ END!")
				rw.WriteString("TKS!")
				rw.Flush()
				break
			default:
				log.Fatalln("READ ERROR!")
			}
		}
		fmt.Println(str)
	}
}
