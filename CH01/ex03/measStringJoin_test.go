// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package measStringJoin

import "testing"

func TestRunmeasStringJoin(t *testing.T) {
	withJoin()
	withoutJoin()
}

func BenchmarkWithoutJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutJoin()
	}
}

func BenchmarkWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withJoin()
	}
}
