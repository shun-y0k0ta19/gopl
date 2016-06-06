// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import "io"

//LimitedReader is a same struct as io.LimitedReader
type LimitedReader struct {
	r io.Reader
	n int64
}

//LimitReader is a same function as  io.LimitReader
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func (l *LimitedReader) Read(p []byte) (int, error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if l.n < int64(len(p)) {
		p = p[0:l.n]
	}
	n, err := l.r.Read(p)
	l.n -= int64(n)
	return n, err
}
