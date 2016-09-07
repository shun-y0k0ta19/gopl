// Copyright © 2016 "Shun Yokota" All rights reserved

package sexpr

import (
	"bufio"
	"os"
)

func ExampleStreamEncode() {
	enc := NewEncoder(os.Stdout)
	enc.Encode(0.0123)
	enc.Encode(5 + 3.3i)
	enc.Encode(make(chan<- string))
	enc.Encode(encode)
	enc.Encode(bufio.NewScanner(os.Stdin))
	type hogeif interface {
	}
	type ifStruct struct {
		hogeif
	}
	hg := ifStruct{[]int{1, 2, 3}}
	enc.Encode(hg)
	//Output:
	//0.0123
	//#C(5 3.3)
	//(chan<- string)
	//(func(*bytes.Buffer, reflect.Value) error)
	//((r ("*os.File" ((file ((fd 0) (name "/dev/stdin") (dirinfo nil)))))) (split (bufio.SplitFunc)) (maxTokenSize 65536) (token ()) (buf ()) (start 0) (end 0) (err nil) (empties 0) (scanCalled nil) (done nil))
	//((hogeif ("[]int" (1 2 3))))
}
