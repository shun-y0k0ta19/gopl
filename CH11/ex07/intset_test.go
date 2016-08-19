// Copyright © 2016 "Shun Yokota" All rights reserved

package intsettest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var seed = time.Now().UTC().UnixNano()
var rng = rand.New(rand.NewSource(seed))

const (
	addNum = 200
	NUM    = 200
)

func BenchmarkMapAdd100(b *testing.B) {
	benchmarkMapAdd(b, 100)
}
func BenchmarkMapAdd10000(b *testing.B) {
	benchmarkMapAdd(b, 10000)
}
func BenchmarkMapAdd1000000(b *testing.B) {
	benchmarkMapAdd(b, 1000000)
}
func BenchmarkMapAdd100000000(b *testing.B) {
	benchmarkMapAdd(b, 100000000)
}

func benchmarkMapAdd(b *testing.B, size int) {
	x := make(MapIntSet)
	for i := 0; i < b.N; i++ {
		x.Add(rng.Intn(size))
	}
}

func BenchmarkMapHas100(b *testing.B) {
	benchmarkMapHas(b, 100)
}
func BenchmarkMapHas10000(b *testing.B) {
	benchmarkMapHas(b, 10000)
}
func BenchmarkMapHas1000000(b *testing.B) {
	benchmarkMapHas(b, 1000000)
}
func BenchmarkMapHas100000000(b *testing.B) {
	benchmarkMapHas(b, 100000000)
}

func setMapIntSet(x MapIntSet, size int) {
	for i := 0; i < addNum; i++ {
		x.Add(rng.Intn(size))
	}
}

func benchmarkMapHas(b *testing.B, size int) {
	x := make(MapIntSet)
	setMapIntSet(x, size)
	for i := 0; i < b.N; i++ {
		x.Has(rng.Intn(size))
	}
}

func BenchmarkMapUnionWith100(b *testing.B) {
	benchmarkMapUnionWith(b, 100)
}

func BenchmarkMapUnionWith10000(b *testing.B) {
	benchmarkMapUnionWith(b, 10000)
}

func BenchmarkMapUnionWith1000000(b *testing.B) {
	benchmarkMapUnionWith(b, 1000000)
}

func BenchmarkMapUnionWith100000000(b *testing.B) {
	benchmarkMapUnionWith(b, 100000000)
}

func benchmarkMapUnionWith(b *testing.B, size int) {
	x := make(MapIntSet)
	y := make(MapIntSet)
	setMapIntSet(x, size)
	setMapIntSet(y, size)
	for i := 0; i < b.N; i++ {
		x.UnionWith(y)
	}
}

func BenchmarkAdd100(b *testing.B) {
	benchmarkAdd(b, 100)
}
func BenchmarkAdd10000(b *testing.B) {
	benchmarkAdd(b, 10000)
}
func BenchmarkAdd1000000(b *testing.B) {
	benchmarkAdd(b, 1000000)
}
func BenchmarkAdd100000000(b *testing.B) {
	benchmarkAdd(b, 100000000)
}

func benchmarkAdd(b *testing.B, size int) {
	var x IntSet
	for i := 0; i < b.N; i++ {
		x.Add(rng.Intn(size))
	}
}

func BenchmarkHas100(b *testing.B) {
	benchmarkHas(b, 100)
}
func BenchmarkHas10000(b *testing.B) {
	benchmarkHas(b, 10000)
}
func BenchmarkHas1000000(b *testing.B) {
	benchmarkHas(b, 1000000)
}
func BenchmarkHas100000000(b *testing.B) {
	benchmarkHas(b, 100000000)
}

func setIntSet(x *IntSet, size int) {
	for i := 0; i < addNum; i++ {
		x.Add(rng.Intn(size))
	}
}

func benchmarkHas(b *testing.B, size int) {
	var x IntSet
	setIntSet(&x, size)
	for i := 0; i < b.N; i++ {
		x.Has(rng.Intn(size))
	}
}

func BenchmarkUnionWith100(b *testing.B) {
	benchmarkUnionWith(b, 100)
}

func BenchmarkUnionWith10000(b *testing.B) {
	benchmarkUnionWith(b, 10000)
}

func BenchmarkUnionWith1000000(b *testing.B) {
	benchmarkUnionWith(b, 1000000)
}

func BenchmarkUnionWith100000000(b *testing.B) {
	benchmarkUnionWith(b, 100000000)
}

func benchmarkUnionWith(b *testing.B, size int) {
	var x IntSet
	var y IntSet
	setIntSet(&x, size)
	setIntSet(&y, size)
	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}

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
