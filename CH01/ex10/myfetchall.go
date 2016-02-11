// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

//result
/*
y0k0ta19:ex10 y0k0ta19$ gor myfetchall.go gopl.io golang.org godoc.org
0.47s     7839  http://golang.org
0.85s     6368  http://godoc.org
1.73s     4146  http://gopl.io
1.73s elapsed
y0k0ta19:ex10 y0k0ta19$ gor myfetchall.go gopl.io golang.org godoc.org
0.29s     4146  http://gopl.io
0.38s     7839  http://golang.org
0.80s     6368  http://godoc.org
0.80s elapsed

*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	resfile, err := os.Create("result.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "myfetchall: %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		chres := <-ch
		fmt.Println(chres) // receive from channel chan type
		fmt.Fprintln(resfile, chres)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	resfile.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
