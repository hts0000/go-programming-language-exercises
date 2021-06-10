package main

import (
	"fmt"
	"strings"
)

// 练习 3.12
// 编写一个函数，判断两个字符串是否是是相互打乱的
// 也就是说它们有着相同的字符，但是对应不同的顺序

func main() {
	s1, s2 := "111666", "166666"
	fmt.Println(charEq1(s1, s2))
	fmt.Println(charEq2(s1, s2))
}

// O(n^2)
func charEq1(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	b := 0
	for _, s1Ch := range s1 {
		for i, s2Ch := range s2 {
			// 如果相等，说明当前s1的字符能够在s2中找到
			if s1Ch == s2Ch {
				b++
				// 剔除s2中已匹配的字符
				s2 = s2[:i] + s2[i+1:]
				break
			}
		}
	}
	if b == len(s1) {
		return true
	}
	return false
}

func charEq2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	b := 0
	for _, s1Ch := range s1 {
		i := strings.LastIndex(s2, string(s1Ch))
		if i != -1 {
			b++
			s2 = s2[:i] + s2[i+1:]
		}
	}
	if b == len(s1) {
		return true
	}
	return false
}
