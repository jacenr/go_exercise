package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/rzcTree/go_exercise/protobuf"
	// "io"
	// "io/ioutil"
	// "bytes"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("connect error.")
	}

	scan := bufio.NewScanner(conn)

	scan.Split(ScanBytes)
	scan.Scan()
	adb_m := scan.Bytes()
	// fmt.Println(adb_m)
	adb := &pb.AddrBook{}
	err = proto.Unmarshal(adb_m, adb)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(adb)
	conn.Close()
}

func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	return 1, data, nil
}
