// Copyright © 2016 "Shun Yokota" All rights reserved

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt" // register PNG decoder
	"log"
	"os"
)

func main() {
	var supportFlag bool
	var formatFlag string
	flag.BoolVar(&supportFlag, "s", false, "List support format.")
	flag.StringVar(&formatFlag, "f", "zip", "Specify archive format.")
	flag.Parse()
	if supportFlag {
		fmt.Println("zip")
		fmt.Println("tar")
	}
	fi := NewArchiveReader("./test.zip")
	fmt.Println(fi)
}

//NewArchiveReader returns zip file reader.
func NewArchiveReader(name string) os.FileInfo {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	fi, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	sys := fi.Sys()
	fmt.Println(sys)
	return fi
}
