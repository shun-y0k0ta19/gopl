// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import "fmt"

func main() {
	s := "abcだよaaΣ𡢽!！"
	fmt.Println(s)
	bs := []byte(s)
	utf8Reverse(bs)
	fmt.Println(string(bs))
}

func utf8Reverse(s []byte) {
	for i := 0; i < len(s); i++ {
		size := utfCharSize(s[i])
		if size > 1 {
			reverse(s[i : i+size])
			i += size - 1
		}
	}
	reverse(s)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func utfCharSize(b byte) int {
	if b <= 0x7F {
		return 1
	}
	if b <= 0xDF {
		return 2
	}
	if b <= 0xEF {
		return 3
	}
	if b <= 0xF7 {
		return 4
	}
	return -1
}
