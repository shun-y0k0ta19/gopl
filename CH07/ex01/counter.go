// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
	"fmt"
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

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	var wc WordCounter
	wc.Write([]byte("fads adfad    dsa fa fafda kjhl"))
	fmt.Println(wc) // "5", = len("hello")

	var lc LineCounter
	lc.Write([]byte(" j \n  saf\r\nfsdfaf"))
	fmt.Println(lc) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

}
