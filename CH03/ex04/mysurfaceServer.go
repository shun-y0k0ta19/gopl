// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var width, height = 900, 600                          // canvas size in pixels
var cells = 100                                       // number of grid cells
var xyrange = float64(30.0)                           // axis ranges (-xyrange..+xyrange)
var xyscale = float64(width) / 3.0 / float64(xyrange) // pixels per x or y unit
var zscale = float64(height) * 0.2                    // pixels per z unit
var angle = math.Pi / 6                               // angle of x, y axes (=30°)

//return z value normalized -1.0 to 1.0
type zf func(float64, float64) float64

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, r)
}

func surface(out io.Writer, r *http.Request) {

	query := r.URL.Query()

	for qname, qvalue := range query {
		var err error
		switch qname {
		case "width":
			width, err = strconv.Atoi(qvalue[0])
		case "height":
			height, err = strconv.Atoi(qvalue[0])
		case "lcolor":
		case "hcolor":
		default:
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "lissajousServer: %v\n", err)
		}
	}

}

func outputSVG(z zf) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, za, isNaNa := corner(i+1, j, z)
			bx, by, zb, isNaNb := corner(i, j, z)
			cx, cy, zc, isNaNc := corner(i, j+1, z)
			dx, dy, zd, isNaNd := corner(i+1, j+1, z)
			if isNaNa || isNaNb || isNaNc || isNaNd {
				continue
			}
			aveZ := int((za + zb + zc + zd) / 4 * 255)
			color := (aveZ << 16) + (255 - aveZ)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%06X'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f zf) (sx float64, sy float64, z float64, isNaN bool) {
	// Find point (x,y) at corner of cell (i,j).
	x, y := xy(i, j)
	// Compute surface height z.
	z = f(x, y)

	if math.IsNaN(z) {
		isNaN = true
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = float64(width/2) + (x-y)*cos30*xyscale
	sy = float64(height/2) + (x+y)*sin30*xyscale - z*zscale

	//(z+1)/2 shift normalized z (-1.0 to 1.0) to (0.0 to 1.0)
	return sx, sy, (z + 1) / 2, isNaN
}

func xy(i, j int) (x, y float64) {
	x = xyrange * (float64(i/cells) - 0.5)
	y = xyrange * (float64(j/cells) - 0.5)
	return x, y
}
