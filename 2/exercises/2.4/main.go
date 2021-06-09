package main

import "fmt"

// 练习 2.4
// 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。
// 比较和查表算法的性能差异。

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	a := PopCount1(911999999)
	fmt.Println(a)
	a = PopCount2(911999999)
	fmt.Println(a)
	a = PopCount3(911999999 - ^911999999 - 1)
	fmt.Println(a)

	// _ = PopCount1(911999999)
	// _ = PopCount2(911999999)
	// _ = PopCount3(911999999 - ^911999999 - 1)
}

// 统计x二进制时有多少个1
func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		// 每次loop，将x的最右bit位 & 1，判断最右bit位是否为1
		count += int(x >> uint(i) & 1)
		// fmt.Println("###", count, uint(i)&1, x>>uint(i)&1)
	}
	return count
}

// 递归实现
func PopCount3(x uint64) int {
	var c int = int(x & 1)
	if x == 1 {
		return 1
	}
	return c + PopCount3(x>>1)
}
