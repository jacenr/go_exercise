package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	// "time"
)

func main() {
	// file := os.Args[1]
	// fi, _ := os.Lstat(file)
	// fmt.Println(fi.Name())
	// fmt.Println(fi.Size())
	// fmt.Println(fi.ModTime())
	// fmt.Println(fi.Mode())
	// fmt.Println(fi.IsDir())
	// fmt.Println(fi.Sys())

	// file1 := "SrcFile"
	f, err := os.Open("SrcFile")
	if err != nil {
		fmt.Println(err)
	}
	// NEWREADER
	// readerF1 := bufio.NewReader(f)

	// NEWSCANNER
	strs := []string{}
	scannerF1 := bufio.NewScanner(f)
	scannerF1.Split(bufio.ScanLines)
	for scannerF1.Scan() {
		strs = append(strs, scannerF1.Text())
	}
	fmt.Println(len(strs))
	for _, txt := range strs {
		fmt.Printf("%s\n", txt)
	}
	fmt.Println(strs[0] == strs[2])
	dirContent, dErr := ioutil.ReadDir("test")
	if dErr != nil {
		fmt.Println(dErr)
	}
	for _, fi := range dirContent {
		fmt.Println(fi.Name())
	}
}
