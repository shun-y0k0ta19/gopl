// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import "time"

import "fmt"

func main() {
	tick := time.NewTicker(1 * time.Second)
	ch := make(chan int)
	go messanger(ch)
	go messanger(ch)
	ch <- 0
	for {
		select {
		case <-tick.C:
			fmt.Printf("rally: %d\n", <-ch)
			ch <- 0
		}
	}
}

func messanger(ch chan int) {
	for {
		msg := <-ch
		msg++
		ch <- msg
	}
}
