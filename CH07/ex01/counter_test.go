// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import "testing"

func TestWordCounter(t *testing.T) {
	const (
		s1wc = 2
		s2wc = 4
	)
	var wc WordCounter
	s1 := []byte("abc abc")
	wc.Write(s1)
	if wc != s1wc {
		t.Errorf("wc.Write(s1) value: %d, expected: %d ", wc, s1wc)
	}

	wc = 0
	s2 := []byte("abc   abc  あ　\tfad   ")
	wc.Write(s2)
	if wc != s2wc {
		t.Errorf("wc.Write(s2) value: %d, expected: %d ", wc, s2wc)
	}
}

func TestLineCounter(t *testing.T) {
	const (
		s1lc = 2
		s2lc = 4
	)
	var lc LineCounter
	s1 := []byte("abc\nabc")
	lc.Write(s1)
	if lc != s1lc {
		t.Errorf("lc.Write(s1) value: %d, expected: %d ", lc, s1lc)
	}

	lc = 0
	s2 := []byte("abc \n  abc \r\n あ　\t\nfad   ")
	lc.Write(s2)
	if lc != s2lc {
		t.Errorf("lc.Write(s2) value: %d, expected: %d ", lc, s2lc)
	}
}
