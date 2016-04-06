// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"testing"
)

//test1
func TestTopologicalOrder(t *testing.T) {
	orderMap := make(map[string]int)
	order, err := topoSort(prereqs)
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
		orderMap[course] = i
	}
	if err != nil {
		t.Errorf("Detect circuration")
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

func TestCircuration(t *testing.T) {
	var circurationPrereqs1 = map[string]strSet{
		"algorithms":     {"data structures": false},
		"calculus":       {"linear algebra": false},
		"linear algebra": {"calculus": false},

		"data structures": {"discrete math": false},
		"discrete math":   {"intro to programming": false},
	}
	var circurationPrereqs2 = map[string]strSet{
		"courseA": {"courseB": false},
		"courseB": {"courseC": false},
		"courseC": {"courseA": false},
	}
	var circurationPrereqs3 = map[string]strSet{
		"courseA": {"courseB": false, "courseC": false},
		"courseC": {"courseA": false},
	}
	if _, err := topoSort(circurationPrereqs1); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Errorf("cannot detect circuration at circurationPrereqs1")
	}
	if _, err := topoSort(circurationPrereqs2); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Errorf("cannot detect circuration at circurationPrereqs2")
	}
	if _, err := topoSort(circurationPrereqs3); err != nil {
		fmt.Println(err.Error())
	} else {
		t.Errorf("cannot detect circuration at circurationPrereqs3")
	}
}

func topoOrderTest(t *testing.T) {
	orderMap := make(map[string]int)
	if order, err := topoSort(prereqs); err != nil {
		t.Errorf("Detect circuration")
	} else {
		for i, course := range order {
			orderMap[course] = i
		}
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
