package main

import "fmt"

func main() {
	var f func(int, int) int
	f = func(b int, c int) int {
		return b + c
	}
	list := []int{0, 1}
	list = calc(list, f)
	f = func(b int, _ int) int {
		return b + 1
	}
	list = calc(list, f)
	fmt.Println(list[len(list)-1])
}

func calc(list []int, f func(int, int) int) []int {
	n := len(list)
	list = append(list, f(list[n-1], list[n-2]))
	return list
}
