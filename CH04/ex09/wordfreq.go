// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		wordfreq(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				log.Fatal(err)
			}
			wordfreq(f, counts, arg)
		}
	}
	for w, c := range counts {
		fmt.Printf("%s\t%d\n", w, c)
	}
}

func wordfreq(in *os.File, counts map[string]int, filename string) {
	input := bufio.NewScanner(in)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}
