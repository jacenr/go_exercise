package main

import (
	"fmt"
	"net"
)

func main() {
	ip1 := net.ParseIP("192.168.1.1") // [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 192, 168, 1, 1], 16字节, 0xff == 255
	fmt.Println(ip1)
	t1 := []byte("192.168.1.1")
	fmt.Println(t1)
	ip2 := net.IP(t1) // ['1','9','2','.','1','6','8','.','1','.','1']
	fmt.Println(ip2)
	for _, i := range ip2 {
		fmt.Println(i)
	}
	fmt.Println()
	fmt.Println(len(ip2))
	fmt.Println()
	for _, i := range ip1 {
		fmt.Println(i)
	}
	// fmt.Printf("%T\n", ip2)
	fmt.Println()
	fmt.Println(len(ip1))
}
