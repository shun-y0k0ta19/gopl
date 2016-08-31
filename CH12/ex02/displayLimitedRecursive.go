// Copyright © 2016 "Shun Yokota" All rights reserved

// Package display provides a means to display structured data.
package display

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const recursiveLimit = 10

//Display shows information of variable.
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	displayLimitedRecursive(name, reflect.ValueOf(x), 0)
}

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func displayLimitedRecursive(path string, v reflect.Value, loop int) {
	loop++
	if loop > recursiveLimit {
		fmt.Printf("%s...\n", path)
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			displayLimitedRecursive(fmt.Sprintf("%s[%d]", path, i), v.Index(i), loop)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			displayLimitedRecursive(fieldPath, v.Field(i), loop)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			var elms []string
			switch key.Kind() {
			case reflect.Array:
				for i := 0; i < key.Len(); i++ {
					elms = append(elms, formatAtom(key.Index(i)))
				}
			case reflect.Struct:
				for i := 0; i < key.NumField(); i++ {
					elms = append(elms, fmt.Sprintf("%s: %s", key.Type().Field(i).Name, formatAtom(key.Field(i))))
				}
			default:
				elms = append(elms, formatAtom(key))
			}
			displayLimitedRecursive(fmt.Sprintf("%s[%s]", path,
				strings.Join(elms, ", ")), v.MapIndex(key), loop)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			displayLimitedRecursive(fmt.Sprintf("(*%s)", path), v.Elem(), loop)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			displayLimitedRecursive(path+".value", v.Elem(), loop)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
