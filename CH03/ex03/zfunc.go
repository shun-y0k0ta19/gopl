// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import "math"

const (
	a = 1
	b = 0
	c = 50
    scale = 1.5
)

func sq(x float64) float64 {
	return x * x
}


func sinrrZ(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)

	res := math.Sin(r) / r
	return res
}

func eggboxZ(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 6
}

func mogulZ(x, y float64) float64 {
	//return (math.Exp(-sq(x-b)/c) * math.Exp(-sq(y-b)/c)) / a
	return (math.Sin(x) * math.Sin(y)) / 4
}



func saddleZ(x, y float64) float64 {
	return (-sq(x/xyrange*scale) + sq(y/xyrange*scale))
}
