package main

import (
	"fmt"
	// "io"
	"log"
	"net"
	"strings"
	"time"
	// "os"
	// "ioutil"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	// for {
	// 	_, err := io.WriteString(c, time.Now().Format("2006 01.02 15:04:05\n"))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	time.Sleep(1 * time.Second)
	// }
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, s string, t time.Duration) {
	fmt.Fprintf(c, "\t%s\n", strings.ToUpper(s))
	time.Sleep(t)
	fmt.Fprintf(c, "\t%s\n", strings.Title(s))
	time.Sleep(t)
	fmt.Fprintf(c, "\t%s\n", strings.ToLower(s))
}
