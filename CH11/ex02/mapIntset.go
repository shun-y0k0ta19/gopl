// Copyright © 2016 "Shun Yokota" All rights reserved

package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet map[int]bool

// Has reports whether the set contains the non-negative value x.
func (s MapIntSet) Has(x int) bool {
	_, ok := s[x]
	return ok
}

// Add adds the non-negative value x to the set.
func (s MapIntSet) Add(x int) {
	s[x] = true
}

// UnionWith sets s to the union of s and t.
func (s MapIntSet) UnionWith(t MapIntSet) {
	for n := range t {
		s[n] = true
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s MapIntSet) String() string {
	var words []int
	for key := range s {
		words = append(words, key)
	}
	sort.Ints(words)
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, word := range words {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", word)
	}
	buf.WriteByte('}')
	return buf.String()
}
