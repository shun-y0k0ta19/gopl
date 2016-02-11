// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

//result
/*
go test -bench .
PASS
BenchmarkWithoutJoin-4	      10	 132607999 ns/op
BenchmarkWithJoin-4   	    5000	    296768 ns/op
ok  	golang_training/CH01/ex03	3.132s
*/

package measStringJoin

import "strings"

var ss []string

func init() {
	for i := 0; i < 20000; i++ {
		ss = append(ss, "a")
	}
}

func withoutJoin() string {
	s, sep := "", ""
	for _, arg := range ss {
		s += sep + arg
		sep = " "
	}
	return s
}

func withJoin() string {
	s := strings.Join(ss, " ")
	return s
}
