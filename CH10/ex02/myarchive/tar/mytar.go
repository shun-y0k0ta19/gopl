// Copyright © 2016 "Shun Yokota" All rights reserved

package tar

import (
	"archive/tar"
	"golang_training/CH10/ex02/myarchive"
	"os"
)

type reader struct {
	tr *tar.Reader
}

func (r *reader) Next() (fileName string, err error) {
	h, err := r.tr.Next()
	if err != nil {
		return "", err
	}
	return h.Name, nil
}

func (r *reader) Read(b []byte) (n int, err error) {
	return r.tr.Read(b)
}

func init() {
	myarchive.RegisterFormat("tar", "ustar", 0x101, newReader)
}

func newReader(f *os.File) (myarchive.Reader, error) {
	tr := tar.NewReader(f)
	return &reader{tr}, nil
}
