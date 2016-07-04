// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import "testing"

func BenchmarkMymandelbrotPall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mymandelbrotPall()
	}
}

func BenchmarkMymandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mymandelbrot()
	}
}
