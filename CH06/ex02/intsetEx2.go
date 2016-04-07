// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds the non-negative value x to the set.
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//Len returns number of elements.
func (s *IntSet) Len() int {
	var count int
	for _, w := range s.words {
		if w != 0 {
			count += popCount(w)
		}
	}
	return count
}

//Remove remove x from this IntSet.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.words[word] == 0 {
		return
	}
	s.words[word] &= ^(1 << bit)
}

//Clear remove all elements.
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

//Copy returns copy of this set.
func (s *IntSet) Copy() *IntSet {
	var newWords IntSet
	newWords.words = make([]uint64, len(s.words))
	copy(newWords.words, s.words)
	return &newWords
}

func popCount(x uint64) int {
	count := 0
	for i := uint64(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}
	return count
}

func main() {
	var s IntSet
	s.Add(1)
	s.Add(4)
	s.Add(12)
	s.Add(100)
	var ss IntSet
	ss.AddAll(1, 4, 12, 100)
	fmt.Printf("s.Len: %d\n", s.Len())
	cp := *s.Copy()
	fmt.Printf("s is %v\n", s.String())
	fmt.Printf("ss is %v\n", ss.String())
	fmt.Printf("cp is %v\n", cp.String())
	s.Remove(4)
	fmt.Printf("s is %v\n", s.String())
	fmt.Printf("cp is %v\n", cp.String())
	fmt.Printf("s has 4: %v\n", s.Has(4))
	fmt.Printf("cp has 4: %v\n", cp.Has(4))
	s.Clear()
	s.Len()
	fmt.Printf("s.Len after Clear(): %d\n", s.Len())

}
