package main

import "fmt"

// 练习 4.7
// 修改reverse函数用于原地反转UTF-8编码的[]byte
// 是否可以不用分配额外的内存？

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Printf("%p\n", s)
	reverse(s)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}

// 没有改变切片长度、容量、底层数组，因此不需要返回新切片
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
