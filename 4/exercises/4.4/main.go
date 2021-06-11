package main

// 练习 4.4
// 编写一个rotate函数，通过一次循环完成旋转

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 = rotateRight(s1, 4)
	fmt.Println(s1)

	s2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 = rotateLeft(s2, 10)
	fmt.Println(s2)
}

// 右旋
func rotateRight(s []int, n int) []int {
	if len(s) == 0 {
		fmt.Println("error: empty slice.")
		return nil
	} else if n > len(s) || n < 0 {
		fmt.Println("error: out of range.")
		return nil
	}

	s = append(s[n:], s[:n]...)
	return s
}

// 左旋
func rotateLeft(s []int, n int) []int {
	if len(s) == 0 {
		fmt.Println("error: empty slice.")
		return nil
	} else if n > len(s) || n < 0 {
		fmt.Println("error: out of range.")
		return nil
	}

	// n%len(s)
	n = len(s) - n
	s = append(s[n:], s[:n]...)
	return s
}
