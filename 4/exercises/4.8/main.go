package main

// 练习 4.8
// 修改charcount程序，使用unicode.IsLetter等相关的函数,
// 统计字母、数字等Unicode中不同的字符类别

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	numners := 0
	letters := 0
	others := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount2: %v\n", err)
			os.Exit(1)
		}
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsNumber(r):
			numners++
		default:
			others++
		}
	}
	fmt.Print("numbers\tletters\tothers\n")
	fmt.Printf("%d\t%d\t%d\n", numners, letters, others)
}
