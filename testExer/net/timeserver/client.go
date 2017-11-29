package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	for {
		// b := []byte{}
		b := make([]byte, 512) // 缓存大小不会自动调整，要么进一步处理，要么使用bufio中的方法.
		n, err := conn.Read(b)
		if err != nil {
			log.Fatalln(err)
		}
		conn.Write([]byte("TKS!"))
		fmt.Println(string(b[:n]))
	}
}
