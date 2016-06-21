// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

//Reader implements io.Reader
type Reader string

func main() {
	url := "https://golang.org"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	s := fmt.Sprint(resp.Body)
	html.Parse(NewReader(s))
}

//NewReader is return io.Reader using s
func NewReader(s string) *Reader {
	r := Reader(s)
	return &r
}

func (r *Reader) Read(p []byte) (int, error) {
	*r = Reader(p)
	return len(p), nil
}
