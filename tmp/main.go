package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/rzcTree/go_exercise/protobuf"
	"io/ioutil"
	"log"
)

func main() {

	p := pb.Person{
		Name:  "Tom",
		Id:    123456,
		Email: "abcdef@abc.com",
		Phones: []*pb.Person_PhoneNumb{
			&pb.Person_PhoneNumb{
				Number: "021123456",
				Type:   pb.Person_HOME,
			},
			&pb.Person_PhoneNumb{
				Number: "13532546355",
				Type:   pb.Person_MOBILE,
			},
		},
	}

	adb := &pb.AddrBook{[]*pb.Person{&p}}

	fmt.Println(adb)

	adb_pb, err := proto.Marshal(adb)
	if err != nil {
		log.Fatalln("Marshal error.")
	}
	if err := ioutil.WriteFile("adb_pb_data", adb_pb, 0644); err != nil {
		log.Fatalln("Write file error.")
	}

	adb_pb_new, err := ioutil.ReadFile("adb_pb_data")
	if err != nil {
		log.Fatalln("Read file err")
	}

	adb_new := &pb.AddrBook{}
	if err := proto.Unmarshal(adb_pb_new, adb_new); err != nil {
		log.Fatalln("Unmarshal error.")
	}
	fmt.Println(adb_new)

	// 临时测试变量
	// tmp_a := struct {
	// 	fd1 string
	// 	fd2 int32
	// }{
	// 	fd1: "test",
	// 	fd2: 123456,
	// }
	// fmt.Println(tmp_a)

	// 测试导入是否成功
	// a := pb.Person_PhoneType(11)
	// fmt.Printf("%T\n%v\n", a, a)
}
