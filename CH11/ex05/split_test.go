// Copyright © 2016 "Shun Yokota" All rights reserved

package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"a b c", " ", 3},
		{"a:b:c", " ", 1},
		{":a:::b::c:", ":", 8},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q %q) returned %d words, want %d", test.s, test.sep, got, test.want)
		}

	}
}
