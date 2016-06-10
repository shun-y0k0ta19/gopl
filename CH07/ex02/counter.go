// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
	"bytes"
)

//ByteCounter is number of byte count
type ByteCounter int

//WordCounter is number of word count
type WordCounter int

//LineCounter is number of line count
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*wc++
	}
	return int(*wc), nil
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	for sc.Scan() {
		*lc++
	}
	return int(*lc), nil
}
