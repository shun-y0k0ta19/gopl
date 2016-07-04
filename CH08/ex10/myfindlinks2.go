// Copyright © 2016 "Shun Yokota" All rights reserved

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

type listInfo struct {
	list  []string
	layer int
}

func crawl(url string, cancel chan struct{}) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := Extract(url, cancel)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	depth := flag.Int("depth", 0, "define crawl depth.\n-depth=0 means no limit to crawling depth.")
	flag.Parse()
	worklist := make(chan listInfo)
	cancel := make(chan struct{})
	var n int // number of pending sends to worklist
	// Start with the command-line arguments.
	n++
	go func() {
		var li listInfo
		li.list = flag.Args()
		worklist <- li
	}()
	go func() {
		sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
		for sc.Scan() {
			if sc.Text() == "q" {
				close(cancel)
			}
		}
	}()
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	var i int
	for ; n > 0; n-- {
		li := <-worklist
		if li.layer >= *depth {
			continue
		}
		i++
		for _, link := range li.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, layer int) {
					var li listInfo
					li.list = crawl(link, cancel)
					li.layer = layer + 1
					worklist <- li
				}(link, li.layer)
			}
		}
	}
}

//!-
