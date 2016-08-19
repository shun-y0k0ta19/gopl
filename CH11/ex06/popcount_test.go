// Copyright © 2016 "Shun Yokota" All rights reserved

package popcountbench

import "testing"

//const max = 0xFFFFFFFFFFFFFFFF
const max = 0xFF

func benchmark(f func(uint64) int, num uint64, loop int) {
	for i := 0; i < loop; i++ {
		f(num)
	}
}

func benchmarkPopcount(b *testing.B, num uint64, loop int) {
	for i := 0; i < b.N; i++ {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
		benchmark(PopCount, num, loop)
	}
}

func benchmarkPopcountByShift(b *testing.B, num uint64, loop int) {
	for i := 0; i < b.N; i++ {
		benchmark(PopCountByShift, num, loop)
	}
}

func benchmarkPopcountByClear(b *testing.B, num uint64, loop int) {
	for i := 0; i < b.N; i++ {
		benchmark(PopCountByClear, num, loop)
	}
}

func BenchmarkPopcountF1L1(b *testing.B) {
	benchmarkPopcount(b, 0xF, 1)
}

func BenchmarkPopcountByShiftF1L1(b *testing.B) {
	benchmarkPopcountByShift(b, 0xF, 1)
}

func BenchmarkPopcountByClearF1L1(b *testing.B) {
	benchmarkPopcountByClear(b, 0xF, 1)
}

func BenchmarkPopcountF1L10(b *testing.B) {
	benchmarkPopcount(b, 0xF, 10)
}

func BenchmarkPopcountByShiftF1L10(b *testing.B) {
	benchmarkPopcountByShift(b, 0xF, 10)
}

func BenchmarkPopcountByClearF1L10(b *testing.B) {
	benchmarkPopcountByClear(b, 0xF, 10)
}

func BenchmarkPopcountF1L100(b *testing.B) {
	benchmarkPopcount(b, 0xF, 100)
}

func BenchmarkPopcountByShiftF1L100(b *testing.B) {
	benchmarkPopcountByShift(b, 0xF, 100)
}

func BenchmarkPopcountByClearF1L100(b *testing.B) {
	benchmarkPopcountByClear(b, 0xF, 100)
}

func BenchmarkPopcountF1L10000(b *testing.B) {
	benchmarkPopcount(b, 0xF, 10000)
}

func BenchmarkPopcountByShiftF1L10000(b *testing.B) {
	benchmarkPopcountByShift(b, 0xF, 10000)
}

func BenchmarkPopcountByClearF1L10000(b *testing.B) {
	benchmarkPopcountByClear(b, 0xF, 10000)
}

func BenchmarkPopcountF16L10(b *testing.B) {
	benchmarkPopcount(b, 0xFFFFFFFFFFFFFFFF, 10)
}

func BenchmarkPopcountByShiftF16L10(b *testing.B) {
	benchmarkPopcountByShift(b, 0xFFFFFFFFFFFFFFF, 10)
}

func BenchmarkPopcountByClearF16L10(b *testing.B) {
	benchmarkPopcountByClear(b, 0xFFFFFFFFFFFFFFFF, 10)
}

func BenchmarkPopcountF16L10000(b *testing.B) {
	benchmarkPopcount(b, 0xFFFFFFFFFFFFFFFF, 10000)
}

func BenchmarkPopcountByShiftF16L10000(b *testing.B) {
	benchmarkPopcountByShift(b, 0xFFFFFFFFFFFFFFF, 10000)
}

func BenchmarkPopcountByClearF16L10000(b *testing.B) {
	benchmarkPopcountByClear(b, 0xFFFFFFFFFFFFFFFF, 10000)
}
