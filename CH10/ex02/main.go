// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bytes"
	"flag"
	"fmt"
	"golang_training/CH10/ex02/myarchive"
	_ "golang_training/CH10/ex02/myarchive/tar"
	_ "golang_training/CH10/ex02/myarchive/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var supportFlag bool
	flag.BoolVar(&supportFlag, "s", false, "List support format.")
	flag.Parse()
	if supportFlag {
		fmt.Println("zip")
		fmt.Println("tar")
	}
	for _, path := range flag.Args() {
		fmt.Printf("Open... %s\n", path)
		r, err := myarchive.NewArchiveReader(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		for {
			name, err := r.Next()
			if err != nil {
				break
			}
			fmt.Printf("Next: %s\n", name)
			buf := new(bytes.Buffer)
			if _, err = io.Copy(buf, r); err != nil {
				log.Fatalln(err)
			}

			if err = ioutil.WriteFile("dc_"+path[2:]+"_"+name, buf.Bytes(), 0755); err != nil {
				log.Fatal(err)
			}
		}
	}
}
