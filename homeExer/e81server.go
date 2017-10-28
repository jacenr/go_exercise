package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ny := flag.String("NewYork", "localhost:8080", "Please give the NewYork addr:port.")
	ld := flag.String("LonDon", "localhost:8081", "Please give the LonDon addr:port.")
	ty := flag.String("Toyo", "localhost:8082", "Please give the Toyo addr:port.")
	flag.Parse()
	addrports := []string{}
	addrports = append(addrports, *ny)
	addrports = append(addrports, *ld)
	addrports = append(addrports, *ty)
	fmt.Println(addrports)
	ch := make(chan string, 3)
	for _, addrport := range addrports {
		fmt.Println(addrport + "listened.")
		go netlisten(addrport, ch)
	}
	// for {
	// }
	for s := range ch {
		fmt.Println(s)
	}
}

func netlisten(adder string, ch chan string) {
	defer func() { ch <- "done" }()
	fmt.Println("ok")
	lsn, err := net.Listen("tcp", adder)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lsn.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go respclient(conn, adder)
	}
}

func respclient(conn net.Conn, adder string) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, adder+" "+time.Now().Format("2006 01 02 15:04:05\n"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}
