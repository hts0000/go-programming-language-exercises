package main

import (
	"fmt"
	"sync"
)

var (
	loadPcOnce sync.Once
	pc         [256]byte
)

func init() {
	loadPcOnce.Do(loadPc)
	// fmt.Println(pc)
}

func main() {
	fmt.Println(PopCount(7))
}

func loadPc() {
	// 生成0~255，每个数二进制中1的个数
	for i := range pc {
		//fmt.Println(pc[i], pc[i/2], byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 统计x二进制时有多少个1
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
