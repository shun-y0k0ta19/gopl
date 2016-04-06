// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"testing"
)

//test1
func TestTopologicalOrder(t *testing.T) {
	orderMap := make(map[string]int)
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
		orderMap[course] = i
	}

	for course, topoOrder := range orderMap {
		fmt.Printf("\ncourse: %s, topoOrder: %d\n", course, topoOrder)
		if prereq, ok := prereqs[course]; ok {
			for precourse := range prereq {
				fmt.Printf("precouse: %s, preCourseOrder: %d\n", precourse, orderMap[precourse])
				if orderMap[precourse] > topoOrder {
					t.Errorf("%s is after %s\n", course, prereq)
				}
			}
		}
	}
}

//test2
func Test100TimesRepeatTopologicalOrder(t *testing.T) {
	for i := 0; i < 100; i++ {
		topoOrderTest(t)
	}
	fmt.Println("done 100 times topoOrderTest")
}

func topoOrderTest(t *testing.T) {
	orderMap := make(map[string]int)
	for i, course := range topoSort(prereqs) {
		orderMap[course] = i
	}

	for course, topoOrder := range orderMap {
		if prereq, ok := prereqs[course]; ok {
			for precourse := range prereq {
				if orderMap[precourse] > topoOrder {
					t.Errorf("%s is after %s\n", course, prereq)
				}
			}
		}
	}
}
