// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
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
	var adsum int
	for {
		ad, token, err := bufio.ScanWords(p[adsum:], true)
		if err != nil {
			return 0, err
		}
		if len(token) > 0 {
			*wc++
		}

		adsum += ad
		if len(p[adsum:]) == 0 {
			return int(*wc), nil
		}
	}
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	var adsum int
	for {
		ad, token, err := bufio.ScanLines(p[adsum:], true)
		if err != nil {
			return 0, err
		}
		if len(token) > 0 {
			*lc++
		}

		adsum += ad
		if len(p[adsum:]) == 0 {
			return int(*lc), nil
		}
	}
}
