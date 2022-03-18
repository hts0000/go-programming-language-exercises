package main

import "fmt"

// 练习 4.5
// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作
// 原地完成————不改变切片底层数组

func main() {
	s := []string{"0", "0", "2", "2", "2", "4", "6", "6", "8", "8"}
	fmt.Printf("%p\n", &s)
	// n表示去重后切片长度
	s = uniqElement(s)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}

func uniqElement(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			fmt.Println(s)
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
