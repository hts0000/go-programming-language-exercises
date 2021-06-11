package main

import "fmt"

func main() {
	// 声明的是切片
	s := []int{0, 1, 2, 3, 4, 5}
	// 声明的是数组，数组必须指定长度
	// s := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)

	s = []int{0, 1, 2, 3, 4, 5}
	reverseNum(s, 3)
	fmt.Println(s)
}

// 没有改变切片长度、容量、底层数组，因此不需要返回新切片
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 没有改变切片长度、容量、底层数组，因此不需要返回新切片
func reverseNum(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}
