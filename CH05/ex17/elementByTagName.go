// Copyright © 2016 "Shun Yokota" All rights reserved
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse: %v", err)
	}

	nodes := ElementByTagName(doc, "h1", "h2", "h3")
	for _, n := range nodes {
		fmt.Printf("%v\n", n.Data)
	}
	nodes = ElementByTagName(doc, "div", "link", "html")
	for _, n := range nodes {
		fmt.Printf("%v\n", n.Data)
	}
}

//ElementByTagName is extracting element by tag name.
func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var elements []*html.Node
	var findElement func(n *html.Node)
	findElement = func(n *html.Node) {
		//fmt.Printf("n.Data: %s\n", n.Data)
		for _, tag := range name {
			if n.Data == tag {
				elements = append(elements, n)
				return
			}
		}
		return
	}
	forEachNode(doc, findElement, nil)

	return elements

}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && n.FirstChild != nil {
		post(n)
	}
}
