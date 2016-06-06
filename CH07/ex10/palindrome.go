// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"sort"
)

type sortable []rune

func (s sortable) Len() int {
	return len(s)
}
func (s sortable) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s1 := sortable("abc")
	fmt.Println(palindrome(s1))
	s2 := sortable("abcba")
	fmt.Println(palindrome(s2))
	s3 := sortable("abccba")
	fmt.Println(palindrome(s3))
}

func palindrome(s sort.Interface) bool {
	last := s.Len() - 1
	for i := 0; i < last-i; i++ {
		if !equals(s, i, last-i) {
			return false
		}
	}
	return true
}

func equals(s sort.Interface, i, j int) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}
