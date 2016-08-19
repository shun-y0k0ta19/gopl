// Copyright © 2016 "Shun Yokota" All rights reserved

package zip

import (
	"archive/zip"
	"golang_training/CH10/ex02/myarchive"
	"io"
	"log"
	"os"
)

type reader struct {
	zr *zip.Reader
	i  int
	cr io.Reader
}

func (r *reader) Next() (fileName string, err error) {
	r.i++
	if r.i >= len(r.zr.File) {
		return "", io.EOF
	}
	f := r.zr.File[r.i]
	rc, err := f.Open()
	if err != nil {
		log.Fatalln(err)
	}
	r.cr = rc
	return r.zr.File[r.i].Name, nil
}

func (r *reader) Read(b []byte) (n int, err error) {
	return r.cr.Read(b)
}

func init() {
	myarchive.RegisterFormat("zip", "PK", 0, newReader)
}

func newReader(f *os.File) (myarchive.Reader, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	zr, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return nil, err
	}

	return &reader{zr: zr, i: -1}, nil
}
