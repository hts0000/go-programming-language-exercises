package main

import (
	"bytes"
	"fmt"
)

//练习6.1
// 为bit数组实现下面这些方法
// func (*IntSet) Len() int      // return the number of elements
// func (*IntSet) Remove(x int)  // remove x from the set
// func (*IntSet) Clear()        // remove all elements from the set
// func (*IntSet) Copy() *IntSet // return a copy of the set

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func main() {
	var x, y IntSet
	// x.words[0]=0b01
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

	fmt.Println(x.Len(), y.Len()) // "4 2"

	// fmt.Println(x.words)
	// fmt.Println(bitCount(4398046511618)) // 3
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	// word: 获取第几个64能存下
	// bit: 获取存放在第几个bit位
	// 以1为例，在words数组中，元素类型为uint64，可以存下64个bit位
	// 1/64=0，words只需要一个元素即可存下，换句话说，存储小于64的数只需要一个元素
	// 1%64=1，s.words[0]的第一个bit位置为1，表示存储了1这个数
	// 以144位例，144/64=2，144%64=16
	// 则将s.words[2]的第16个bit位置1，表示存储了144这个数
	word, bit := x/64, uint(x%64)
	// 切片扩容，直到能存下x
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
			// 判断当前bit位是否不为0
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j) // 根据存储的规则，将数还原
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 官方解法:https://github.com/ray-g/gopl/blob/8cd4330890081305af7eff2ac09f1a821f99c9de/ch06/ex6.01/intset.go
// 返回这个数中有几个bit位为1
func bitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

// return the number of elements
func (s *IntSet) Len() int {
	len := 0

	// 官方解法
	// for _, word := range s.words {
	// 	len += bitCount(word)
	// }

	// 解法1
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			// 判断当前bit位是否不为0
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len

	// 解法2
	// return len(strings.Split(s.String(), " ")) // can work!
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	// 找到这个数对应的数组下标和bit位，将bit位置0即为删除这个数
	word, bit := x/64, uint(x%64)
	s.words[word] -= 1 << bit
}

// remove all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	intSet := new(IntSet)
	intSet.words = make([]uint64, len(s.words))
	copy(intSet.words, s.words)
	return intSet
}
