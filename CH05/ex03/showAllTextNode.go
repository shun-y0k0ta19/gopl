// Copyright © 2016 "Shun Yokota"" All rights reserved

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
		//s = ""
		fmt.Println(s)
		//fmt.Print(s)
	}
}

func recursiveVisit(links []string, n *html.Node) []string {
	if n != nil {
		/*
			fmt.Printf("n.Type : %d, TextNode :%d\n", n.Type, html.TextNode)
			fmt.Printf("n.Namespace : %s\n", n.Namespace)
			fmt.Printf("n.DataAtom : %x\n", n.DataAtom)
			fmt.Printf("n.Attribute: %v\n", n.Attr)
			fmt.Printf("n.Data: %v\n", n.Data)
		*/
		if n.Type == html.TextNode {
			if n.Parent.Data != "script" && n.Parent.Data != "style" {
				links = append(links, n.Data)
			}
		}
		links = recursiveVisit(recursiveVisit(links, n.NextSibling), n.FirstChild)
	}

	return links
}
