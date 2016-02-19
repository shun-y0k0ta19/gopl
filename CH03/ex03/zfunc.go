// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import "math"

const (
	sinrrScale  = 1
	eggboxScale = 0.5
	mogulsScale = 0.5
	saddleScale = 1.0
)

func sq(x float64) float64 {
	return x * x
}

func sinrrZ(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	z := math.Sin(r) / r
	return z * sinrrScale
}

func eggboxZ(x, y float64) float64 {
	z := (math.Sin(x) + math.Sin(y)) / 2 //normalize -1.0 to 1.0
	return z * eggboxScale
}

func mogulsZ(x, y float64) float64 {
	z := (math.Sin(x) * math.Sin(y)) //normalize -1.0 to 1.0
	return z * mogulsScale
}

func saddleZ(x, y float64) float64 {
	z := -sq(x/(xyrange*0.5)) + sq(y/(xyrange*0.5)) //normalize -1.0 to 1.0
	return z * saddleScale
}
