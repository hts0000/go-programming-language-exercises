package main

import (
	"bytes"
	"fmt"
)

// 练习 6.3
// (*IntSet).UnionWith会用|操作符计算两个集合的并集
// 我们再为IntSet实现另外的几个函数
// IntersectWith(交集：元素在A集合B集合均出现)
// DifferenceWith(差集：元素出现在A集合，未出现在B集合)
// SymmetricDifference(并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{1 144}"

	x.Add(42)
	x.IntersectWith(&y)
	fmt.Println(x.String()) // "{42}"

	x.Add(100)
	x.SymmetricDifference(&y)
	fmt.Println(x.String()) // "{9 100}"
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

// IntersectWith(交集：元素在A集合B集合均出现)
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	s.words = s.words[:len(t.words)]
}

// DifferenceWith(差集：元素出现在A集合，未出现在B集合)
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		// ^: 异或运算，值不同则为1，相同为0
		// B           A
		// 100011010 ^ 110101001 = 010110011
		// 010110011 & 110101001 = 010100001 // 正好是A中有而B中没有的
		if i < len(s.words) {
			s.words[i] &= s.words[i] ^ tword
		}
	}
}

// SymmetricDifference(并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A)
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}
}
