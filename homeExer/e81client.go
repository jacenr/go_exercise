package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addrport := flag.String("addrport", "localhost:8080", "Please give the addr:port to be connected.")
	flag.Parse()
	fmt.Println(*addrport)
	conn, err := net.Dial("tcp", *addrport)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustcopy(os.Stdout, conn)
}

func mustcopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
