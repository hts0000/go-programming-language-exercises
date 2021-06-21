package main

import "fmt"

func main() {
	strs := []string{"123", "", "321"}

	// 错误用法，没有更新strs，println输出的是一个新切片
	// 书中原话：
	// 更新slice变量不仅对调用append函数是必要的，
	// 实际上对应任何可能导致长度、容量或底层数组变化的操作都是必要的

	// fmt.Println(nonempty(strs))
	// fmt.Println(nonempty2(strs))
	// fmt.Println()

	// 正确用法
	strs = nonempty(strs)
	fmt.Println(strs)
	strs = nonempty2(strs)
	fmt.Println(strs)
}

func nonempty(strs []string) []string {
	i := 0
	for _, s := range strs {
		if s != "" {
			strs[i] = s
			i++
		}
	}
	// 返回了一个新切片，切片本身地址和strs地址不一样
	return strs[:i]
}

func nonempty2(strs []string) []string {
	// out依然指向strs底层数组
	out := strs[:0]
	// fmt.Println(out)
	// fmt.Println("out == nil", out == nil) // false
	for _, s := range strs {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
