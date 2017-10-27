package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
一次性读取整改文件实例:
import (
	"io/ioutil"
)
data,err := ioutil.ReadFile("filename")

strings包中有Split和Join函数
_, line := range strings.Split(string(data), "\n")
*/

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	if len(files) == 0 {
		countlines(os.Stdin, counts)
	}
	for _, arg := range files {

		// 打开文件的方式: os.open
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			continue
		}
		countlines(f, counts)
		f.Close()
	}
	for line, n := range counts {
		fmt.Printf("%d,%s\n", n, line)
	}
}

// 需要参数 文件f, map类型counts
func countlines(f *os.File, counts map[string]int) {

	// 获取输入的方式
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
