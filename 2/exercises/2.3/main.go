package main

var pc [256]byte

func init() {
	// 生成0~255，每个数二进制中1的个数
	for i := range pc {
		// fmt.Println(pc[i], pc[i/2], byte(i&1))
		pc[i] = pc[i/2] + byte(i&1)
	}
	// fmt.Println(pc)
}

func main() {
	// start := time.Now()
	_ = PopCount1(999999)
	// end := time.Since(start).Nanoseconds()
	// fmt.Println(end, "ns")

	// start = time.Now()
	_ = PopCount2(999999)
	// end = time.Since(start).Nanoseconds()
	// fmt.Println(end, "ns")
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
	var c byte
	for i := 0; i < 8; i++ {
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}
