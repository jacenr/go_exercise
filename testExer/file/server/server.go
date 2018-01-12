package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
)

// 可以通过[]byte和[]string的方式传输文件
// 可以分片或者整体传输文件
// 但考虑到内存和文件的大小不定, 最好分片传输, 就像在gosync中演示的那样

type mg struct {
	List []string
}

func main() {
	f, err := os.Open("tmp")
	fmt.Println(err)
	defer f.Close()
	frd := bufio.NewReader(f)
	slist := []string{}
	for {
		line, err := frd.ReadString('\n')
		slist = append(slist, line)
		if err == io.EOF {
			break
		}
	}
	m := mg{slist}
	lst, _ := net.Listen("tcp", "127.0.0.1:8999")
	for {
		conn, err := lst.Accept()
		fmt.Println(err)
		cnwt := bufio.NewWriter(conn)
		enc := gob.NewEncoder(cnwt)
		err = enc.Encode(m)
		fmt.Println(err)
		cnwt.Flush()
	}
}
