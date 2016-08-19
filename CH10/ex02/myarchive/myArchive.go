// Copyright © 2016 "Shun Yokota" All rights reserved

package myarchive

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type archiveFormat struct {
	name, magic string
	offset      int
	newReader   func(*os.File) (Reader, error)
}

//Reader is the interface wraps methods to read archive.
type Reader interface {
	Next() (fileName string, err error)
	Read(b []byte) (n int, err error)
}

var formats []archiveFormat

//NewArchiveReader returns zip file reader.
func NewArchiveReader(name string) (Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	af := findArchiveFormat(f)
	if af.newReader == nil {
		return nil, nil
	}
	f.Seek(0, os.SEEK_SET)
	return af.newReader(f)
}

//RegisterFormat registers format of archive.
func RegisterFormat(name, magic string, offset int, reader func(*os.File) (Reader, error)) {
	formats = append(formats, archiveFormat{name, magic, offset, reader})
}

func findArchiveFormat(f *os.File) archiveFormat {
	r := bufio.NewReader(f)
	for _, af := range formats {
		h, err := r.Peek(af.offset + len(af.magic))
		if err == io.EOF {
			continue
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if equals(h[af.offset:], []byte(af.magic)) {
			return af
		}
	}
	return archiveFormat{}
}

func equals(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i, b := range b1 {
		if b2[i] != b {
			return false
		}
	}
	return true
}
