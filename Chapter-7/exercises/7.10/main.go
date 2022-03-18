package main

import (
	"fmt"
	"sort"
)

// 练习 7.10
// sort.Interface类型也可以适用在其它地方
// 编写一个IsPalindrome(s sort.Interface) bool函数
// 表明序列s是否是回文序列，换句话说反向排序不会改变这个序列
// 假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等

type Str string

func (s Str) Len() int           { return len(s) }
func (s Str) Less(i, j int) bool { return s[i] == s[j] }
func (s Str) Swap(i, j int)      {}

func main() {
	var s Str = "aaasddddddaddsaaa"
	fmt.Println(s)
	if IsPalindrome(s) {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是回文")
	}
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < s.Len()/2; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			return false
		}
	}
	return true
}
