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

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for word > 0 {
			count += int(word & 1)
			word = word >> 1
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		fmt.Printf("%s: %d\n", "NO THIS VALUE.", x)
		return
	}
	word, bit := x/64, uint(x%64)
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	// var d *IntSet = &IntSet{[]uint64{}}
	words := make([]uint64, len(s.words))
	var d *IntSet = &IntSet{words}
	// fmt.Println(d.words == nil)
	copy(d.words, s.words)
	return d
}

func main() {
	is1 := IntSet{[]uint64{1, 2, 3}}
	fmt.Println(is1)
	fmt.Println(is1.String())
	is3 := is1.Copy()
	fmt.Println("is3: " + is3.String())
	// fmt.Println(is1.Len())
	is1.Remove(300)
	fmt.Println(is1.String())
	is1.Remove(128)
	fmt.Println(is1.String())
	is1.Clear()
	fmt.Println(is1.String())
	var is2 IntSet
	// is2.Add(0)
	is2.Add(65)
	// is2.Add(128)
	// is2.Add(129)
	fmt.Println(is2)
	fmt.Println(is2.String())
	fmt.Println(is2.Len())
}
