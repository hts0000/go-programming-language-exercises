package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1(args []string) (s string) {
	var sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return
}

func echo2(args []string) (s string) {
	var sep string
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return
}

func echo3(args []string) (s string) {
	s = strings.Join(args[1:], " ")
	return
}

func main() {
	args := os.Args
	var start time.Time
	var end int64

	start = time.Now()
	echo1(args)
	end = time.Since(start).Microseconds()
	fmt.Println("echo1 use time:", end, "us")

	start = time.Now()
	echo2(args)
	end = time.Since(start).Microseconds()
	fmt.Println("echo2 use time:", end, "us")

	start = time.Now()
	_ = echo3(args)
	end = time.Since(start).Nanoseconds()
	fmt.Println("echo3 use time:", end, "ns")
}
