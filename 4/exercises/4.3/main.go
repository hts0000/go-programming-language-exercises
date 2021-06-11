package main

// 练习 4.3
// 重写reverse函数，使用数组指针代替slice

import (
	"fmt"
)

const arrayLen = 10

func main() {
	array := [arrayLen]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&array)
	fmt.Println(array)
}

// go中的数组作为参数是值拷贝，需要传入数组指针
func reverse(array *[arrayLen]int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
