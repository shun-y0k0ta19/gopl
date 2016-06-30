// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		timezoneAndURL := strings.Split(arg, "=")
		timezone := timezoneAndURL[0]
		url := timezoneAndURL[1]
		fmt.Printf("arg: %s\n", arg)
		fmt.Printf("timeZone: %s\n", timezone)
		fmt.Printf("url: %s\n", url)
	}
	/*
		conn, err := net.Dial("tcp", "localhost:8010")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		mustCopy(os.Stdout, conn)
	*/
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
