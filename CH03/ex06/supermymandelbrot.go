// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 2048, 2048
)

func main() {

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
	png.Encode(os.Stdout, superSampling(img)) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
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

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func superSampling(source image.Image) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
	for py := 0; py < img.Rect.Dy(); py++ {
		for px := 0; px < img.Rect.Dx(); px++ {
			img.Set(px, py, average(source, px*2, py*2))
		}
	}
	return img
}

func average(source image.Image, px int, py int) color.Color {
	var colors [4]color.Color
	colors[0] = source.At(px, py)
	colors[1] = source.At(px+1, py)
	colors[2] = source.At(px, py+1)
	colors[3] = source.At(px+1, py+1)
	var sr, sg, sb uint32
	for _, c := range colors {
		//fmt.Println(n)
		r, g, b, _ := c.RGBA()
		sr += r
		sg += g
		sb += b
	}
	return color.RGBA{
		uint8(sr / 4),
		uint8(sg / 4),
		uint8(sb / 4),
		255,
	}
}
