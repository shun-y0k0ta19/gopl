// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "世kai界Σ　　　　だ 𡢽   dayo"
	bs := []byte(s)
	res := replaceSpace(bs)

	fmt.Printf("%v\n", []byte(s))
	fmt.Println(res)
	fmt.Println(s)
	fmt.Println(string(res))
}

func replaceSpace(bytes []byte) []byte {
	for i := 0; i < len(bytes); {
		size := utfCharSize(bytes[i])
		if isUnicodeSpace(bytes[i:]) {
			if isUnicodeSpace(bytes[i+size:]) {
				bytes[i] = 0x20
				for j := 1; j < size; j++ {
					bytes = remove(bytes, i+1)
				}
				bytes = removeUTF8(bytes, i+1)
				continue
			}

		}
		i += size
	}
	return bytes
}

func isUnicodeSpace(bytes []byte) bool {
	r, _ := extractUTF8(bytes)
	return unicode.IsSpace(r)
}

func extractUTF8(bytes []byte) (rune, int) {
	size := utfCharSize(bytes[0])
	return utf8.DecodeRune(bytes[0:size])
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

func removeUTF8(slice []byte, i int) []byte {
	size := utfCharSize(slice[i])
	for j := 0; j < size; j++ {
		slice = remove(slice, i)
	}
	return slice
}

func remove(slice []byte, i int) []byte {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
