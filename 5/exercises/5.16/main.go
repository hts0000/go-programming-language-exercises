package main

import (
	"fmt"
	"strings"
)

// 练习5.16
// 编写多参数版本的strings.Join

func main() {
	fmt.Println(strJoin("*", "1", "2", "asaa", "sdfa"))
	fmt.Println(strJoin("*", "aaa"))
	fmt.Println(strings.Join([]string{"1", "2", "asaa", "sdfa"}, "*"))
	fmt.Println(strings.Join([]string{"aaa"}, "*"))
}

func strJoin(sep string, strs ...string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	comstr := ""
	for _, str := range strs {
		comstr += str + sep
	}
	return comstr
}
