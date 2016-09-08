// Copyright © 2016 "Shun Yokota" All rights reserved

package myequal2

import (
	"bytes"
	"fmt"
	"testing"

	"gopl.io/ch13/equal"
)

func TestEqual(t *testing.T) {
	one, oneAgain, two := 1, 1, 2

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice CycleSlice
	cycleSlice = append(cycleSlice, cycleSlice)

	ch1, ch2 := make(chan int), make(chan int)
	var ch1ro <-chan int = ch1

	type mystring string

	var iface1, iface1Again, iface2 interface{} = &one, &oneAgain, &two

	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		// basic types
		{1, 1, true},
		{1, 2, false},   // different values
		{1, 1.0, false}, // different types
		{"foo", "foo", true},
		{"foo", "bar", false},
		{mystring("foo"), "foo", false}, // different types
		// slices
		{[]string{"foo"}, []string{"foo"}, true},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string{}, []string(nil), true},
		// slice cycles
		{cycleSlice, cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3}},
			true,
		},
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3, 4}},
			false,
		},
		{
			map[string][]int{},
			map[string][]int(nil),
			true,
		},
		// pointers
		{&one, &one, true},
		{&one, &two, false},
		{&one, &oneAgain, true},
		{new(bytes.Buffer), new(bytes.Buffer), true},
		// pointer cycles
		{cyclePtr1, cyclePtr1, true},
		{cyclePtr2, cyclePtr2, true},
		{cyclePtr1, cyclePtr2, true}, // they're deeply equal
		// functions
		{(func())(nil), (func())(nil), true},
		{(func())(nil), func() {}, false},
		{func() {}, func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, [...]int{1, 2, 3}, true},
		{[...]int{1, 2, 3}, [...]int{1, 2, 4}, false},
		// channels
		{ch1, ch1, true},
		{ch1, ch2, false},
		{ch1ro, ch1, false}, // NOTE: not equal
		// interfaces
		{&iface1, &iface1, true},
		{&iface1, &iface2, false},
		{&iface1Again, &iface1, true},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}

func TestCycle(t *testing.T) {
	one := 1

	type CyclePtr *CyclePtr
	var cyclePtr1 CyclePtr
	cyclePtr1 = &cyclePtr1

	type CycleSlice []CycleSlice
	var cycleSlice CycleSlice
	cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	var iface1 interface{} = &one

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{1, false},
		{one, false},
		{1.0, false},
		{"foo", false},
		// slices
		{[]string{"foo"}, false},
		// slice cycles
		{cycleSlice, true},
		// maps

		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},

		// pointers
		{&one, false},
		{new(bytes.Buffer), false},
		// pointer cycles
		{cyclePtr1, true},
		// arrays
		{[...]int{1, 2, 3}, false},
		// interfaces
		{&iface1, false},
	} {
		if CheckCircle(test.x) != test.want {
			t.Errorf("CheckCircle(%v) = %t",
				test.x, !test.want)
		}
	}

}

func Example_equal() {
	//!+
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))        // "true"
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))      // "false"
	fmt.Println(Equal([]string(nil), []string{}))             // "true"
	fmt.Println(Equal(map[string]int(nil), map[string]int{})) // "true"
	const (
		a = 1.0
		b = 10.0
		c = 2.0
	)
	var f1 float64
	var f2 float64
	f1 = float64(a / b * c)
	for i := 0; i < 10; i++ {
		f := float64(0.02)
		f2 += f
	}
	fmt.Println(Equal(f1, f2))         // "true"
	fmt.Println(equal.Equal(f1, f2))   // "false"
	fmt.Println(Equal(&f1, &f2))       // "true"
	fmt.Println(equal.Equal(&f1, &f2)) // "false"
	var cmplx1 complex128
	var cmplx2 complex128
	cmplx1 = 1.0 + 1.0i
	for i := 0; i < 10; i++ {
		cmplx := 0.1 + 0.1i
		cmplx2 += cmplx
	}
	fmt.Println(Equal(cmplx1, cmplx2))       // "true"
	fmt.Println(equal.Equal(cmplx1, cmplx2)) // "false"
	// Output:
	// true
	// false
	// true
	// true
	// true
	// false
	// true
	// false
	// true
	// false
}

func Example_equalCycle() {
	//!+cycle
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a)) // "true"
	fmt.Println(Equal(b, b)) // "true"
	fmt.Println(Equal(c, c)) // "true"
	fmt.Println(Equal(a, b)) // "false"
	fmt.Println(Equal(a, c)) // "false"
	//!-cycle

	// Output:
	// true
	// true
	// true
	// false
	// false
}
