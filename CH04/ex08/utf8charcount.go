// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)          // counts of Unicode characters
	classfication := make(map[string]int) //counts of Unicode classification
	var utflen [utf8.UTFMax + 1]int       // count of lengths of UTF-8 encodings
	invalid := 0                          // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsControl(r) {
			classfication["Control"]++
		}
		if unicode.IsDigit(r) {
			classfication["Digit"]++
		}
		if unicode.IsGraphic(r) {
			classfication["Graphic"]++
		}
		if unicode.IsLetter(r) {
			classfication["Letter"]++
		}
		if unicode.IsLower(r) {
			classfication["Lower"]++
		}
		if unicode.IsMark(r) {
			classfication["Mark"]++
		}
		if unicode.IsNumber(r) {
			classfication["Number"]++
		}
		if unicode.IsPrint(r) {
			classfication["Print"]++
		}
		if unicode.IsPunct(r) {
			classfication["Punct"]++
		}
		if unicode.IsSpace(r) {
			classfication["Space"]++
		}
		if unicode.IsSymbol(r) {
			classfication["Symbol"]++
		}
		if unicode.IsTitle(r) {
			classfication["Title"]++
		}
		if unicode.IsUpper(r) {
			classfication["Upper"]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nclassification\tcount\n")
	for cls, n := range classfication {
		fmt.Printf("%s\t%d\n", cls, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
