package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "./home/hts/test.txt"
	fmt.Println(basename(s))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
