// Copyright © 2016 "Shun Yokota" All rights reserved

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 2048, 2048
)

func mymandelbrot() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
}

func mymandelbrotPall() *image.RGBA {
	var wg sync.WaitGroup
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)
			go func(px, py int, y float64) {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
				wg.Done()
			}(px, py, y)
		}
	}
	wg.Wait()
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 1000
	const contrast = 15

	var v complex128
	for i := 0; i < iterations; i++ {
		v = v*v + z
		n := uint8(i)
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				255 - contrast*n*(n%2),
				255 - contrast*n*(n%3),
				255 - contrast*n*(n%5),
				255,
			}
		}
	}
	return color.Black
}

func main() {
	//mymandelbrotPall()
	img := mymandelbrot()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors

}

/*
ゴルーチン5個
BenchmarkMymandelbrotPall-4	       1	4404917893 ns/op
BenchmarkMymandelbrot-4    	       1	4604223087 ns/op

ゴルーチン50個
BenchmarkMymandelbrotPall-4	       1	3803288881 ns/op
BenchmarkMymandelbrot-4    	       1	4575116026 ns/op

ゴルーチン500個
BenchmarkMymandelbrotPall-4	       1	3335388594 ns/op
BenchmarkMymandelbrot-4    	       1	4661580789 ns/op

ゴルーチン5000個
BenchmarkMymandelbrotPall-4	       1	3345505204 ns/op
BenchmarkMymandelbrot-4    	       1	4605279790 ns/op

ゴルーチン50000個
BenchmarkMymandelbrotPall-4	       1	3597591985 ns/op
BenchmarkMymandelbrot-4    	       1	4643917877 ns/op

*/
