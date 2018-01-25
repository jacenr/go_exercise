package main

import (
	"fmt"
	"strings"
)

func main() {
	m := map[string]string{"1": "aagag", "2": "gghre", "3": "ahrtesj", "4": "htsetgre", "5": "ahterjtr", "6": "jrstw3g", "7": "ah5j64w", "8": "gfdshe34", "9": "ahterj43t"}
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
		if strings.HasPrefix(v, "a") {
			delete(m, k)
		}
	}
	fmt.Println(m)
}

// 在range的同时delete map中的key, 可以输出预期的结果
