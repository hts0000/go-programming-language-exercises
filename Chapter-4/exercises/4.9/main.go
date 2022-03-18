package main

import (
	"bufio"
	"fmt"
	"os"
)

// 练习 4.9
// 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率
// 在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数
// 这样可以按单词而不是按行输入

func main() {
	words := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := in.Text()
		words[word]++
	}
	fmt.Println(words)
}
