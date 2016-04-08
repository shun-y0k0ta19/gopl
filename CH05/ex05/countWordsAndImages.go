// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountsWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "CountsWordsAndImages: %s\n", err)
	}
	fmt.Printf("words: %d, images: %d\n", words, images)
	/*
		recursiveList, imageNum := recursiveVisit(nil, 0, doc)

		for _, s := range recursiveList {
			//s = ""
			fmt.Println(s)
			//fmt.Print(s)
		}
	*/
}

//CountsWordsAndImages does an HTTP GET Request for the HTML
//document url and return the number of words and images in it.
func CountsWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(h *html.Node) (words, images int) {
	textList, images := recursiveVisit(nil, 0, h)
	textScanner := bufio.NewScanner(strings.NewReader(strings.Join(textList, " ")))
	textScanner.Split(bufio.ScanWords)
	for textScanner.Scan() {
		words++
	}
	return
}

func recursiveVisit(links []string, images int, n *html.Node) ([]string, int) {
	if n != nil {
		if n.Type == html.TextNode {
			if n.Parent.Data != "script" && n.Parent.Data != "style" {
				links = append(links, n.Data)
			}
		} else if n.Type == html.ElementNode {
			if n.Data == "img" {
				images++
			}
		}
		links, images = recursiveVisit(links, images, n.FirstChild)
		links, images = recursiveVisit(links, images, n.NextSibling)
	}

	return links, images
}
