// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s1 := "hoge"
	s2 := "hoge"
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	fmt.Printf("Difference of %s, %s: %d\n", s1, s2, popCountDiff(c1, c2))

	s1 = "hoge"
	s2 = "huge"
	c1 = sha256.Sum256([]byte(s1))
	c2 = sha256.Sum256([]byte(s2))
	fmt.Printf("Difference of %s, %s: %d\n", s1, s2, popCountDiff(c1, c2))
}

func popCountDiff(c1, c2 [32]byte) int {
	var count int
	for i := range c1 {
		count += PopCount(uint64(c1[i]) ^ uint64(c2[i]))
	}
	return count
}
