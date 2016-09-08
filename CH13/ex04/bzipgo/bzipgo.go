// Copyright © 2016 "Shun Yokota" All rights reserved

package bzipgo

import (
	"io"
	"os/exec"
)

type writer struct {
	cmd *exec.Cmd
	out io.Writer
}

// NewWriter returns new io.WriteCloser for bzip2
func NewWriter(out io.Writer) io.WriteCloser {
	subProcess := exec.Command("bzip2", "-z")
	return &writer{cmd: subProcess, out: out}
}

func (w *writer) Write(data []byte) (int, error) {
	return 0, nil
}

// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {

	return nil
}
