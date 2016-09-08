// Copyright © 2016 "Shun Yokota" All rights reserved

// Package myequal2 provides a deep equivalence relation for arbitrary values.
package myequal2

import (
	"math"
	"math/cmplx"
	"reflect"
	"unsafe"
)

const epsilon = 1e-9

func myEqual(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// ...cycle check omitted (shown later)...

	// cycle check
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	// ...numeric cases omitted for brevity...

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return x.Int() == y.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return x.Uint() == y.Uint()

	case reflect.Float32, reflect.Float64:
		return math.Abs(x.Float()-y.Float()) < epsilon

	case reflect.Complex64, reflect.Complex128:
		return cmplx.Abs(x.Complex()-y.Complex()) < epsilon

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return myEqual(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !myEqual(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	// ...struct and map cases omitted for brevity...

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !myEqual(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !myEqual(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

// Equal reports whether x and y are deeply equal.
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return myEqual(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

// CheckCircle checks circle struct
func CheckCircle(x interface{}) bool {
	seen := make(objInfoSet)
	return checkCircle(reflect.ValueOf(x), seen)
}

func checkCircle(x reflect.Value, seen objInfoSet) bool {
	if !x.IsValid() {
		return false
	}
	if !x.CanAddr() {
	} else {
		objinfo := objInfo{unsafe.Pointer(x.UnsafeAddr()), x.Type()}
		if _, ok := seen[objinfo]; ok {
			return true
		}
		seen[objinfo] = true
	}
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return checkCircle(x.Elem(), copyObjInfoSet(seen))

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if checkCircle(x.Index(i), copyObjInfoSet(seen)) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if checkCircle(x.Field(i), copyObjInfoSet(seen)) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if checkCircle(x.MapIndex(k), copyObjInfoSet(seen)) {
				return true
			}
		}
		return false
	}
	return false
}

func copyObjInfoSet(origin objInfoSet) objInfoSet {
	copy := make(objInfoSet)
	for k, v := range origin {
		copy[k] = v
	}
	return copy
}

type objInfo struct {
	x unsafe.Pointer
	t reflect.Type
}

type objInfoSet map[objInfo]bool

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}
