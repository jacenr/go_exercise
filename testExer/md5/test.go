package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

func main() {
	test1, _ := ioutil.ReadFile("test1.txt")
	test2, _ := ioutil.ReadFile("test2.txt")
	test3, _ := ioutil.ReadFile("test3.txt")
	m1 := md5.Sum(test1)
	m2 := md5.Sum(test2)
	m3 := md5.Sum(test3)
	fmt.Printf("%x\n", m1)
	fmt.Printf("%x\n", m2)
	fmt.Printf("%x\n", m3)
	// 只要文件内容一样，得到的md5是一样的。
}
