// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 900, 600            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 3 / xyrange // pixels per x or y unit
	zscale        = height * 0.2        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

//return z value normalized -1.0 to 1.0
type zf func(float64, float64) float64

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	for _, shape := range os.Args[1:] {
		switch shape {
		case "sinrr":
			outputSVG(sinrrZ)
		case "eggbox":
			outputSVG(eggboxZ)
		case "moguls":
			outputSVG(mogulsZ)
		case "saddle":
			outputSVG(saddleZ)
		default:
			outputSVG(sinrrZ)
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
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale

	//(z+1)/2 shift normalized z (-1.0 to 1.0) to (0.0 to 1.0)
	return sx, sy, (z + 1) / 2, isNaN
}

func xy(i, j int) (x, y float64) {
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)
	return x, y
}
