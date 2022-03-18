package main

import (
	"bytes"
	"fmt"
)

// 练习 3.10
// 编写一个非递归版本的comma函数
// 使用bytes.Buffer代替字符串链接操作

func main() {
	s := "-123456"
	fmt.Println(comma(s))
}

func comma(s string) string {
	// 无意义的判断
	// if len(s) <= 3 {
	// 	return s
	// }
	var buf bytes.Buffer
	for i, v := range s {
		if i%3 == 0 && i != 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
