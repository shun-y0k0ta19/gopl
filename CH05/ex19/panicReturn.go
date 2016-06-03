// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import "fmt"

func main() {
	res := increment(1)
	fmt.Printf("res: %d\n", res)
}

func increment(x int) (y int) {
	defer func() {
		if p := recover(); p == x {
			y = x + 1
		} else {
			panic(p)
		}
	}()
	panic(x)
}
