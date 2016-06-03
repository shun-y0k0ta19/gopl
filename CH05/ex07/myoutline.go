// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := outline(url); err != nil {
			log.Print(err)
		}
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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

var depth int

func startElement(n *html.Node) {
	var attributes []string
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			s := fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
			attributes = append(attributes, s)
		}
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range attributes {
			fmt.Print(attr)
		}
		if n.FirstChild == nil {
			fmt.Print("/")
		}
		fmt.Println(">")
		depth++
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
