// Copyright © 2016 "Shun Yokota" All rights reserved

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

type strSet map[string]bool

var prereqs = map[string]strSet{
	"algorithms": {"data structures": false},
	"calculus":   {"linear algebra": false},

	"compilers": {
		"data structures":       false,
		"formal languages":      false,
		"computer organization": false,
	},

	"data structures":       {"discrete math": false},
	"databases":             {"data structures": false},
	"discrete math":         {"intro to programming": false},
	"formal languages":      {"discrete math": false},
	"networks":              {"operating systems": false},
	"operating systems":     {"data structures": false, "computer organization": false},
	"programming languages": {"data structures": false, "computer organization": false},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]strSet) []string {
	var order []string
	determine := make(strSet)
	seen := make(strSet)
	var visitAll func(items strSet) error

	visitAll = func(items strSet) error {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
				determine[item] = false
			} else if _, ok := determine[item]; !ok {
				return fmt.Errorf("circuration at %s", item)
			}
		}
		return nil
	}

	keys := make(strSet)
	for key := range m {
		keys[key] = false
	}

	visitAll(keys)
	return order
}
