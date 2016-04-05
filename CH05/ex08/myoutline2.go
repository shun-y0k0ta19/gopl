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
	id := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse: %v", err)
	}
	ElementByID(doc, id)
}

//ElementByID return html node that found by ID at first
func ElementByID(doc *html.Node, id string) *html.Node {
	return doc
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	var done bool
	if pre != nil {
		if done = pre(n); done {
			return done
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if done = forEachNode(c, pre, post); done {
			return done
		}
	}

	if post != nil && n.FirstChild != nil {
		if done = post(n); done {
			return done
		}
	}
	return done
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
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
