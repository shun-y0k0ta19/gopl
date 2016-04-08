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

	recursiveList := recursiveVisit(nil, doc)

	for _, s := range recursiveList {
		fmt.Println(s)
	}
}

func recursiveVisit(links []string, n *html.Node) []string {
	if n != nil {
		if n.Type == html.TextNode {
			if n.Parent.Data != "script" && n.Parent.Data != "style" {
				links = append(links, n.Data)
			}
		}
		links = recursiveVisit(recursiveVisit(links, n.NextSibling), n.FirstChild)
	}

	return links
}
