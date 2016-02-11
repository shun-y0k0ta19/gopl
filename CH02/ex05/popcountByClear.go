// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package popcountByClear

//PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	for x != 0 {
		x &= (x - 1)
		count++
	}
	return count
}
