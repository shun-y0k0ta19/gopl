// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package popcount2

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var count int
	for i := uint64(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
