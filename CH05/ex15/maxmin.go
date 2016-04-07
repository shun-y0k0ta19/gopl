// Copyright © 2016 "Shun Yokota" All rights reserved
package main

import "fmt"

func max(vals ...int) int {
	n := len(vals)
	if n <= 0 {
		return 0
	}
	for i, v := range vals[:n-1] {
		if v > vals[i+1] {
			vals[i+1] = v
		}
	}
	return vals[n-1]
}

func min(vals ...int) int {
	n := len(vals)
	if n <= 0 {
		return 0
	}
	for i, v := range vals[:n-1] {
		if v < vals[i+1] {
			vals[i+1] = v
		}
	}
	return vals[n-1]
}

func main() {
	fmt.Println(max())              //  "0"
	fmt.Println(max(3))             //  "3"
	fmt.Println(max(1, 2, 3, 4, 3)) //  "4"

	values := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	fmt.Println(max(values...)) // "5"

	fmt.Println(min())              //  "0"
	fmt.Println(min(3))             //  "3"
	fmt.Println(min(1, 2, 3, 4, 3)) //  "4"

	values = []int{3, 2, 1, 2, 3, 4, 5}
	fmt.Println(min(values...)) // "1"
}
