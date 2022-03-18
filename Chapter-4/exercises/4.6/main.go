package main

// 练习 4.6
// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回

import (
	"fmt"
)

func main() {
	s := "h  e \t	l		lo world   !   "
	fmt.Printf("%p\n", &s)

	s = string(uniqSpace([]byte(s)))

	fmt.Printf("%p\n", &s)
	fmt.Println(s)
}

func uniqSpace(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
			s[i] = ' '
		}
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
