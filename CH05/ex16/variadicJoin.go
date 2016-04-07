// Copyright © 2016 "Shun Yokota" All rights reserved
package main

import "fmt"

func main() {
	url := variadicJoin(" ", "go", "lang", ".org")
	fmt.Println(url)
}

func variadicJoin(sep string, a ...string) string {
	fmt.Println(len(a))
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
