// Copyright © 2016 "Shun Yokota" All rights reserved

package intset

import (
	"fmt"
	"testing"
)

const NUM = 200

func TestAdd(t *testing.T) {
	expected := "{1 9 144}"
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(144)
	if expected != x.String() { // "{1 9 144}"
		t.Errorf("expected: {1 9 144}, Actual: %s\n", x.String())
	}
	y := make(MapIntSet)
	y.Add(1)
	y.Add(144)
	y.Add(9)
	y.Add(144)
	if expected != y.String() { // "{1 9 144}"
		t.Errorf("expected: {1 9 144}, Actual: %s\n", x.String())
	}

}

func TestHas(t *testing.T) {
	var x IntSet
	y := make(MapIntSet)
	for i := 0; i < NUM; i++ {
		x.Add(i)
		y.Add(i)
	}
	for i := 0; i < 2*NUM; i++ {
		xok := x.Has(i)
		yok := y.Has(i)
		if i < NUM {
			if !xok {
				t.Errorf("IntSet does not have %d.\n", i)
			}
			if !yok {
				t.Errorf("MapIntSet does not have %d.\n", i)
			}
		} else {
			if xok {
				t.Errorf("IntSet has %d.\n", i)
			}
			if yok {
				t.Errorf("MapIntSet has %d.\n", i)
			}

		}
	}
}

func TestUnionWith(t *testing.T) {
	expected := map[int]bool{
		1:   true,
		144: true,
		9:   true,
		42:  true,
	}
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.UnionWith(&y)
	for i := 0; i < NUM; i++ {
		_, ok := expected[i]
		if ok != x.Has(i) {
			t.Errorf("IntSet can not calculate union with %d.\n", i)
		}
	}
}

func Example_one() {

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func TestUnionWithForMap(t *testing.T) {
	expected := map[int]bool{
		1:   true,
		144: true,
		9:   true,
		42:  true,
	}
	x := make(MapIntSet)
	y := make(MapIntSet)
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.UnionWith(y)
	for i := 0; i < NUM; i++ {
		_, ok := expected[i]
		if ok != x.Has(i) {
			t.Errorf("IntSet can not calculate union with %d.\n", i)
		}
	}
}
