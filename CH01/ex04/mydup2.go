// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type set map[string]struct{}

var includedFiles = make(map[string]set)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var filenames []string
			filenameSet := includedFiles[line]
			for filename := range filenameSet {
				filenames = append(filenames, filename)
			}
			sort.Strings(filenames)
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(filenames, " "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if includedFiles[line] == nil {
			includedFiles[line] = make(set)
		}
		includedFiles[line][filename] = struct{}{}
	}
}
