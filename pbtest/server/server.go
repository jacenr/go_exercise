package main

import (
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
	person := pb.Person{
		Name:  "jerry",
		Id:    123455,
		Email: "dafdsa@abc.com",
		Phones: []*pb.Person_PhoneNumb{
			&pb.Person_PhoneNumb{
				Number: "1256254635",
				Type:   pb.Person_MOBILE,
			},
			&pb.Person_PhoneNumb{
				Number: "02152515168",
				Type:   pb.Person_HOME,
			},
		},
	}
	adb := &pb.AddrBook{[]*pb.Person{&person}}
	// fmt.Println(adb)
	fmt.Printf("Server: %v\n", adb)
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("listen error.")
	}
	for {
		conn, err := ln.Accept()
		defer conn.Close()
		fmt.Println("Connected.")
		if err != nil {
			log.Fatalln("connect error")
		}
		adb_m, err := proto.Marshal(adb)
		if err != nil {
			log.Fatalln("Marshal error.")
		}
		// adb_m_s := bytes.NewBuffer(adb_m)
		// io.Copy(conn, adb_m_s)

		conn.Write(adb_m)
		// conn.Write([]byte("test"))
		fmt.Println("Send.")
	}

}
