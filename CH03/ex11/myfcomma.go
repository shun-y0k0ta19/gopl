// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(fcomma("-461298"))
	fmt.Println(fcomma("-889223456.324005"))
}

func comma(s string, buf *bytes.Buffer) *bytes.Buffer {
	fmt.Printf("comma :%d\n", len(s))
	for i := 0; i < len(s); i++ {
		if (len(s)-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	fmt.Printf("comma buf :%s\n", buf.String())
	fmt.Printf("comma buf2 :%s\n", buf.String())
	return buf
}

func fcomma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	if strings.HasPrefix(s, "-") {
		buf.WriteByte(s[0])
		s = strings.TrimLeft(s, "-")
	}
	point := strings.Index(s, ".")
	fmt.Println(point)
	if point == -1 {
		buf = *comma(s, &buf)
		fmt.Println(len(buf.String()))
	} else {
		fmt.Printf("pbuf1 :%s\n", buf.String())
		buf = *comma(s[:point], &buf)
		fmt.Printf("pbuf2 :%s\n", buf.String())
		for i := point; i < len(s); i++ {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}
