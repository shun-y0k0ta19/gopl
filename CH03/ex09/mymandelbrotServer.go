// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x, err := strconv.ParseFloat(r.FormValue("x"), 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseFloat(r.FormValue("x"), 64)
		if err != nil {
			log.Fatal(err)
		}
		zoom, err := strconv.ParseFloat(r.FormValue("zoom"), 64)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, draw(x, y, zoom))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func draw(cx float64, cy float64, zoom float64) *image.RGBA {

	var xmin, ymin, xmax, ymax = -zoom, -zoom, +zoom, +zoom
	var width, height = 1024, 1024

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x+cx, y+cy)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	return img
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
