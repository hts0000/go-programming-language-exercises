package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

// 练习 3.11
// 完善comma函数，以支持浮点数处理和一个可选的正负号的处理
// go run .\main.go -n -1234556.1111 -N

var isnegative bool

func main() {
	var s string
	flag.BoolVar(&isnegative, "N", false, "numbers is negative")
	flag.StringVar(&s, "n", "12345.12345", "set numbers")
	flag.Parse()

	fmt.Println(comma(s))
}

func comma(s string) string {
	var buf bytes.Buffer
	var f string
	if isnegative {
		s = s[1:]
		buf.WriteRune('-')
	}

	if strings.LastIndex(s, ".") != -1 {
		f = "." + s[strings.LastIndex(s, ".")+1:]
	}
	s = s[:strings.LastIndex(s, ".")]
	for i, v := range s {
		if i%3 == 0 && i != 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(v)
	}
	buf.WriteString(f)
	return buf.String()
}
