// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if (len(s)-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
