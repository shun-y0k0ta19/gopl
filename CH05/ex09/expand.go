// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "GOPATH is $GOPATH\nGOROOT is $GOROOT\n"
	fmt.Println(expand(s, os.Getenv))
}

func expand(s string, f func(string) string) string {
	textScanner := bufio.NewScanner(strings.NewReader(s))
	textScanner.Split(bufio.ScanWords)
	for textScanner.Scan() {
		word := textScanner.Text()
		if strings.HasPrefix(word, "$") {
			res := fmt.Sprintf("%s", f(strings.TrimLeft(word, "$")))
			s = strings.Replace(s, word, res, -1)
		}
	}
	return s
}
