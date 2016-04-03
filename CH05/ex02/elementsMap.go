// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	elementsMap := make(map[string]int)
	recursiveVisit(elementsMap, doc)
	for eName, count := range elementsMap {
		fmt.Printf("ElementName :%s, Count :%d\n", eName, count)
	}
}

func recursiveVisit(elementsMap map[string]int, n *html.Node) {
	if n != nil {
		if n.Type == html.ElementNode {
			elementsMap[n.Data]++
		}
		recursiveVisit(elementsMap, n.NextSibling)
		recursiveVisit(elementsMap, n.FirstChild)
	}
}
