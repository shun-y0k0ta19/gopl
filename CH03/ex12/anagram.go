// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"strings"
)

func anagram(s1 string, s2 string) bool {
	for _, r := range s2 {
		if !strings.Contains(s1, string(r)) {
			return false
		}
		s1 = strings.Replace(s1, string(r), "", 1)
	}
	if len(s1) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(anagram("abc", "abc"))
	fmt.Println(anagram("abc", "bac"))
	fmt.Println(anagram("aabcd", "baca"))
	fmt.Println(anagram("aabca", "aabcaa"))
	fmt.Println(anagram("aabcab", "aabcaa"))
	fmt.Println(anagram("横田峻", "峻横田"))
}
