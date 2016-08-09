// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"time"
)

//const max = 100000

func main() {
	//for i := 1; i < max; i *= 2 {
	for i := uint(1); ; i *= 2 {
		//fmt.Printf("goroutines: %10d\n", i-1)
		//go func(i uint) {
		start, last := starter(i)
		s := time.Now()
		start <- true
		<-last
		e := time.Now().Sub(s)
		fmt.Printf("goroutin:%10d, time: %d\n", i, e.Nanoseconds())
		//}(i - 1)
	}
}

func starter(pNum uint) (chan bool, chan bool) {
	//channels := make([]chan bool, pNum)
	//lastCh := make(chan bool)
	start := make(chan bool)
	var in, out chan bool
	in = start
	for i := uint(0); i < pNum; i++ {
		out = make(chan bool)
		go pipe(in, out)
		in = out
	}
	/*
		for i := range channels {
			channels[i] = make(chan bool)
		}
		for i := range channels {
			if i < len(channels)-1 {
				go pipe(channels[i], channels[i+1])
			} else {
				go pipe(channels[i], lastCh)
			}
		}
	*/
	//return channels[0], lastCh
	return start, out
}

func pipe(in <-chan bool, out chan<- bool) {
	ch := <-in
	out <- ch
}
