// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import "fmt"

func main() {
	type CycleSlice []CycleSlice
	var cycleSlice CycleSlice
	cycleSlice = append(cycleSlice, cycleSlice)
	cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice
	for i := range cycleSlice {
		fmt.Println(cycleSlice[i])
	}
	fmt.Printf("cycleSice: %v\n", cycleSlice)

	type cycle struct {
		value int
		tail  *cycle
	}
	var c cycle
	c = cycle{42, &c}
}
