package main

import (
	"fmt"
	// "os"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.IsAbs("/test/mygo/src"))
	absPath, _ := filepath.Abs("/mygo/src") // 如果path不以"/"开头，则会把当前目录的绝对路径附加到此路径前面
	fmt.Println(absPath)
}
