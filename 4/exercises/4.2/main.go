package main

// 练习 4.2
// 编写一个程序，默认情况下打印标准输入的SHA256编码
// 并支持通过命令行flag定制，输出SHA384或SHA512哈希算法

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var lens string
	flag.StringVar(&lens, "l", "256", "set SHA length, support: 256, 384, 512.")
	flag.Parse()
	for _, arg := range flag.Args() {
		switch lens {
		case "256":
			fmt.Printf("%x\n", sha256.Sum256([]byte(arg)))
		case "384":
			fmt.Printf("%x\n", sha512.Sum384([]byte(arg)))
		case "512":
			fmt.Printf("%x\n", sha512.Sum512([]byte(arg)))
		default:
			flag.Usage()
		}
	}
}
