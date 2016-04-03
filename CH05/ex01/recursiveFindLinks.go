// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	correctList := visit(nil, doc)
	sort.Strings(correctList)

	recursiveList := recursiveVisit(nil, doc)
	sort.Strings(recursiveList)

	if !reflect.DeepEqual(correctList, recursiveList) {
		fmt.Println("recursiveVisit is not correct!!!")
		for i := range correctList {
			fmt.Printf("visit :%s\nrecursiveVisit :%s\n", correctList[i], recursiveList[i])
		}
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func recursiveVisit(links []string, n *html.Node) []string {
	if n != nil {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		links = recursiveVisit(recursiveVisit(links, n.FirstChild), n.NextSibling)
	}

	return links
}
