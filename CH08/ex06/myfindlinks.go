// Copyright © 2016 "Shun Yokota" All rights reserved

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

type listInfo struct {
	list  []string
	layer int
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
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
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		var li listInfo
		li.list = flag.Args()
		worklist <- li
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	var i int
	for ; n > 0; n-- {
		fmt.Printf("n: %d\n", n)
		li := <-worklist
		if li.layer >= *depth {
			continue
		}
		i++
		fmt.Printf("--------------layer %d------------------------\n", i)
		for _, link := range li.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, layer int) {
					var li listInfo
					li.list = crawl(link)
					li.layer = layer + 1
					worklist <- li
				}(link, li.layer)
			}
		}
	}
}

//!-
