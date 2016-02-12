// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette handler
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w, r)
}

func trimLR(s, sep string) (left string, right string) {
	nameAndParam := strings.Split(s, sep)
	return nameAndParam[0], nameAndParam[1]
}

func lissajous(out io.Writer, r *http.Request) {
	uri := strings.Trim(r.RequestURI, "/?") //Query解析を使う
	varname, param := trimLR(uri, "=")

	cycles := 5   // number of complete x oscillator revolutions
	res := 0.001  // angular resolution
	size := 100   // image canvas covers [-size..+size]
	nframes := 64 // number of animation frames
	delay := 8    // delay between frames in 10ms units

	var err error

	switch varname {
	case "cycles":
		cycles, err = strconv.Atoi(param)
	case "size":
		size, err = strconv.Atoi(param)
	case "nframes":
		nframes, err = strconv.Atoi(param)
	case "delay":
		delay, err = strconv.Atoi(param)
	default:
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajousServer: %v\n", err)
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
