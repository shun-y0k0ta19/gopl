// Copyright © 2016 "Shun Yokota" All rights reserved

package popcountbench

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//PopCountByShift returns the population count (number of set bits) of x.
func PopCountByShift(x uint64) int {
	count := 0
	for i := uint64(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}
	return count
}

//PopCountByClear returns the population count (number of set bits) of x.
func PopCountByClear(x uint64) int {
	count := 0
	for x != 0 {
		x &= (x - 1)
		count++
	}
	return count
}
