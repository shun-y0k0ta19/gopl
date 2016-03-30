// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"sort"

	"golang.org/x/net/html"
)

func TestRecursiveVisit(t *testing.T) {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	correctList := visit(nil, doc)
	sort.Strings(correctList)

	var recursiveList []string
	/*
		recursiveList := recursiveVisit(nil, doc)
		sort.Strings(recursiveList)
	*/
	if reflect.DeepEqual(correctList, correctList) {
		t.Errorf("CorrectList : %v recursiveList : %v", correctList, recursiveList)
	}
	/*
		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	*/
}
