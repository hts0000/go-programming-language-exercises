package main

// 练习 4.1
// 编写一个函数，计算两个SHA256哈希码中不同bit的数目
// （参考2.6.2节的PopCount函数。)

import (
	"fmt"
)

func main() {
	// c1 := sha256.Sum256([]byte("x"))
	// c2 := sha256.Sum256([]byte("X"))

	c1 := [32]byte{'1', '2'}
	c2 := [32]byte{'3', '4'}
	fmt.Printf("%d\n%d\n", c1, c2)
	fmt.Printf("%b\n%b\n", c1, c2)

	fmt.Println("#####", diffBit(c1, c2))
}

func diffBit(c1, c2 [32]byte) uint32 {
	var count uint32
	for i := 0; i < len(c1); i++ {
		fmt.Printf("i: %d, diffBit:\n%8b\n%8b\n", i, c1[i], c2[i])
		// 循环8次，一个byte为8bit
		for j := 0; j < 8; j++ {
			// 判断最右的bit位是否相同，右移刷新位置
			if (c1[i]>>j)&1 != (c2[i]>>j)&1 {
				count++
			}
		}
		fmt.Println("count =", count)
		fmt.Println()
	}
	return count
}
