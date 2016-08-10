// Copyright © 2016 "Shun Yokota" All rights reserved

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	var supportFlag bool
	var formatFlag string
	flag.BoolVar(&supportFlag, "s", false, "List support format.")
	flag.StringVar(&formatFlag, "f", "jpeg", "Specify format to convert.")
	flag.Parse()
	if supportFlag {
		fmt.Println("jpeg")
		fmt.Println("png")
		fmt.Println("gif")
	}
	switch formatFlag {
	case "jpeg":
		if err := toJPEG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "convert to JPEG\n")
	case "png":
		if err := toPNG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "convert to PNG\n")
	case "gif":
		if err := toGIF(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "gif: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "convert to GIF\n")
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return gif.Encode(out, img, &gif.Options{NumColors: 255, Quantizer: nil, Drawer: nil})
}
