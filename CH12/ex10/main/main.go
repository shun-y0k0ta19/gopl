package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Fprintf(os.Stderr, "recover: %v\n", p)
			fmt.Printf("recover: %v\n", p)
		}
	}()
	type intst []int
	//ints := []int{1, 2, 3}
	ints := intst{1, 2, 3}
	v := reflect.ValueOf(ints)
	fmt.Println(v)
	fmt.Println(v.Kind())
	fmt.Println(v.Type())
	fmt.Println(v.Type().Elem())
	fmt.Println(v.Type().Elem().Elem())
	fmt.Println(reflect.New(v.Type().Elem()))
	fmt.Println(reflect.New(v.Type().Elem()).Elem())
	k := reflect.Kind(1)
	fmt.Println(k.String())

}
