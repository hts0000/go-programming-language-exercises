package main

import "fmt"

func main() {
	strack := []int{1, 2, 3}
	strack = strackPush(4, strack)
	fmt.Println(strack)
	p := strackPop(strack)
	fmt.Println(p)
}

func strackPush(v int, strack []int) []int {
	return append(strack, v)
}

func strackPop(strack []int) int {
	p := strack[len(strack)-1]
	strack = strack[:len(strack)-1]
	return p
}
