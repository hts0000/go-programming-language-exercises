package main

import "fmt"

// 练习 4.5
// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作
// 原地完成————不改变切片长度、容量、底层数组

func main() {
	s := []int{0, 0, 2, 2, 2, 4, 6, 6, 8, 8}
	// n表示去重后切片长度
	s = uniqElement(s)
	fmt.Println(s)
}

func uniqElement(s []int) []int {
	l := len(s)
	for i := 0; i < l; i++ {
		j := i + 1
		for ; j < len(s); j++ {
			if s[j] != s[i] {
				break
			}
		}
		fmt.Println("con")
		s = append(s[:i+1], s[j:]...)
	}
	return s
}
