// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"testing"
)

func TestNormalize(t *testing.T) {
	testNornalize(t, sinrrZ)
	testNornalize(t, eggboxZ)
	testNornalize(t, mogulsZ)
	testNornalize(t, saddleZ)
}

func testNornalize(t *testing.T, zfunc zf) {
	fv := reflect.ValueOf(zfunc)
	//get function name
	zfuncName := fmt.Sprint(runtime.FuncForPC(fv.Pointer()).Name())
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y := xy(i, j)
			z := zfunc(x, y)
			if math.Abs(z) > 1.0 {
				t.Errorf("z over 1.0! i,j=(%d,%d) x,y=(%.3f,%.3f) z=%.3f@%v", i, j, x, y, z, zfuncName)
			}
		}
	}
}
