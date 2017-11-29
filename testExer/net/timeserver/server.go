package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	lst, err := net.Listen("tcp", "127.0.0.1:6000")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := lst.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		conn.Write([]byte(time.Now().String()))
		// b := []byte{}
		b := make([]byte, 512) // 缓存大小不会自动调整，要么进一步处理(for循环)，要么使用bufio中的方法.
		n, err := conn.Read(b)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(b[:n]))
		time.Sleep(1 * time.Second)
	}
}
