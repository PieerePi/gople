// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

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

// Len returns the number of elements.
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if word&(1<<j) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove clears the non-negative value x from the set.
func (s *IntSet) Remove(x int) {
	if x >= len(s.words)*64 {
		return
	}
	word, bit := x/64, uint(x%64)
	s.words[word] &^= (1 << bit)
}

// Clear clears all elements from the set.
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy returns a copy of the set.
func (s *IntSet) Copy() *IntSet {
	snew := &IntSet{words: make([]uint64, len(s.words))}
	copy(snew.words, s.words)
	return snew
}

// AddAll adds the non-negative values to the set.
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// Intersection sets s to the intersection of s and t.
func (s *IntSet) Intersection(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
}

// Difference sets s to the difference of s and t.
func (s *IntSet) Difference(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems returns all elements of the set.
func (s *IntSet) Elems() []int {
	var elems = make([]int, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if word&(1<<j) != 0 {
				elems = append(elems, i*64+int(j))
			}
		}
	}
	return elems
}

// Elems2 returns all elements of the set.
func (s *IntSet) Elems2() []int {
	var elems = make([]int, s.Len())
	count := 0
	for i := range s.words {
		if s.words[i] == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if s.words[i]&(1<<j) != 0 {
				elems[count] = i*64 + int(j)
				count++
			}
		}
	}
	return elems
}

// Elems3 returns all elements of the set.
func (s *IntSet) Elems3() []int {
	var elems = make([]int, 0, s.Len())
	for i := range s.words {
		if s.words[i] == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if s.words[i]&(1<<j) != 0 {
				elems = append(elems, i*64+int(j))
			}
		}
	}
	return elems
}

// Elems4 returns all elements of the set.
func (s *IntSet) Elems4() []int {
	var elems = make([]int, s.Len())
	count := 0
	for i := 0; i < len(s.words); i++ {
		if s.words[i] == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if s.words[i]&(1<<j) != 0 {
				elems[count] = i*64 + int(j)
				count++
			}
		}
	}
	return elems
}

// Elems5 returns all elements of the set.
func (s *IntSet) Elems5() []int {
	var elems = make([]int, s.Len())
	count := 0
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if word&(1<<j) != 0 {
				elems[count] = i*64 + int(j)
				count++
			}
		}
	}
	return elems
}

// Elems6 returns all elements of the set.
// c style loop
func (s *IntSet) Elems6() []int {
	var elems []int
	for i := 0; i < len(s.words); i++ {
		if s.words[i] == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if s.words[i]&(1<<j) != 0 {
				elems = append(elems, i*64+int(j))
			}
		}
	}
	return elems
}

// Elems7 returns all elements of the set.
// use index of range
func (s *IntSet) Elems7() []int {
	var elems []int
	for i := range s.words {
		if s.words[i] == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if s.words[i]&(1<<j) != 0 {
				elems = append(elems, i*64+int(j))
			}
		}
	}
	return elems
}

// Elems8 returns all elements of the set.
// use index, value of range
func (s *IntSet) Elems8() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if word&(1<<j) != 0 {
				elems = append(elems, i*64+int(j))
			}
		}
	}
	return elems
}

//!-intset

//!+string

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

//!-string
