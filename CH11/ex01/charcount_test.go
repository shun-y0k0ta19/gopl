// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input   string
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{"./test1.txt", map[rune]int{'a': 1, 'b': 1, 'c': 1}, []int{0, 3}, 0},
	}
	for _, test := range tests {
		f, err := os.Open(test.input)
		if err != nil {
			t.Fatal(err)
		}
		counts, utflen, invalid := charCount(bufio.NewReader(f))
		if len(counts) != len(test.counts) {
			var ess, ass []string
			for c, n := range counts {
				ess = append(ess, fmt.Sprintf("%q\t%d\n", c, n))
			}
			es := strings.Join(ess, "")
			for c, n := range test.counts {
				ass = append(ass, fmt.Sprintf("%q\t%d\n", c, n))
			}
			as := strings.Join(ass, "")
			t.Errorf("\nexpected:\nrune\tcount\n%v\nactual:\nrune\tcount\n%v\n", es, as)
		}
		fmt.Println(utflen)
		fmt.Println(invalid)
	}
}
