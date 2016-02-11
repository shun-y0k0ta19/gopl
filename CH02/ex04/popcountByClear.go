// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package popcountByClear

//PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	for i := uint64(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}
	return count
}
