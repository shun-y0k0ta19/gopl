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
	node, err := ElementByID(doc, id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, attr := range node.Attr {
		fmt.Printf("key: %s, value: %s\n", attr.Key, attr.Val)
	}
}

//ElementByID return html node that found by ID at first
func ElementByID(doc *html.Node, id string) (*html.Node, error) {
	var idElement *html.Node
	var findElement func(n *html.Node) bool
	findElement = func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				idElement = n
				return true
			}
		}
		return false
	}
	forEachNode(doc, findElement, findElement)
	var err error
	if idElement == nil {
		err = fmt.Errorf("id=%s is not found.\n", id)
	}

	return idElement, err
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
