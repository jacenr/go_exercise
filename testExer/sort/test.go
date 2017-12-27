package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	// "io/ioutil"
	"sort"
)

func testsort(fileName string) {
	f, _ := os.Open(fileName)
	rd := bufio.NewReader(f)

	var rdErr error
	var strs []string
	var line string
	for rdErr != io.EOF {
		line, rdErr = rd.ReadString('\n')
		strs = append(strs, line)
	}
	sort.Strings(strs)
	for _, v := range strs {
		fmt.Printf("%s", v)
	}
	fmt.Println(len(strs))
}

func main() {
	testsort("dst271031")
	// 对不到40000行的文本进行排序，排序速度还能够接收。
}
