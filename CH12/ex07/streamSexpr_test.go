// Copyright © 2016 "Shun Yokota" All rights reserved

package sexpr

import "os"

func ExampleStreamEncode() {
	enc := NewEncoder(os.Stdout)
	enc.Encode(0.0123)
	enc.Encode(5 + 3.3i)
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
	//((hogeif ("[]int" (1 2 3))))
}
