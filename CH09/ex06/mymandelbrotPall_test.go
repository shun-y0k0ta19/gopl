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

/*
y0k0ta19:ex06 y0k0ta19$ GOMAXPROCS=1 got -bench .
testing: warning: no tests to run
PASS
BenchmarkMymandelbrotPall	       1	8636706785 ns/op
BenchmarkMymandelbrot    	       1	4666345145 ns/op
ok  	golang_training/CH09/ex06	13.395s
y0k0ta19:ex06 y0k0ta19$ GOMAXPROCS=2 got -bench .
testing: warning: no tests to run
PASS
BenchmarkMymandelbrotPall-2	       1	4295489040 ns/op
BenchmarkMymandelbrot-2    	       1	4602561453 ns/op
ok  	golang_training/CH09/ex06	8.950s
y0k0ta19:ex06 y0k0ta19$ GOMAXPROCS=4 got -bench .
testing: warning: no tests to run
PASS
BenchmarkMymandelbrotPall-4	       1	3216817985 ns/op
BenchmarkMymandelbrot-4    	       1	4585152656 ns/op
ok  	golang_training/CH09/ex06	7.843s
*/
