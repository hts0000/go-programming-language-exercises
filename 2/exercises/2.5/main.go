package main

import "fmt"

// 练习 2.5
// 表达式x&(x-1)用于将x的最低的一个非零的bit位清零
// 使用这个算法重写PopCount函数，然后比较性能

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	a := PopCount2(911999999)
	fmt.Println(a)
	a = PopCount2(911999999)
	fmt.Println(a)
}

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

// 从右开始，不断地将x最右边为1的bit为置0，x=0时停止
// 循环的次数就是x bit位为1的个数
// 第一次循环：1101 & 1100 = 1100
// 第二次循环：1100 & 1011 = 1000
// 第三次循环：1000 & 0111 = 0000
func PopCount2(x uint64) int {
	var c int
	for ; x != 0; x &= x - 1 {
		c++
	}
	return c
}
