// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinksEx: %v\n", err)
		os.Exit(1)
	}

	recursiveList := recursiveVisit(nil, doc)
	sort.Strings(recursiveList)

	for _, s := range recursiveList {
		fmt.Println(s)
	}

}

func recursiveVisit(links []string, n *html.Node) []string {
	if n != nil {
		fmt.Printf("n.Type : %d, TextNode :%d\n", n.Type, html.TextNode)
		fmt.Printf("n.Namespace : %s\n", n.Namespace)
		fmt.Printf("n.DataAtom : %x\n", n.DataAtom)
		fmt.Printf("n.Attribute: %v\n", n.Attr)
		fmt.Printf("n.Data: %v\n", n.Data)

		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "href" || a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
		links = recursiveVisit(recursiveVisit(links, n.FirstChild), n.NextSibling)
	}

	return links
}
