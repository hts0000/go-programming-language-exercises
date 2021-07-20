package main

import (
	"fmt"
	"syscallerr/syscall"
)

func main() {
	var err error = syscall.Erron(2)
	fmt.Println(err.Error())
	fmt.Println(err)
	// err实现了error接口，打印输出时会自动调用Error方法
	// 为什么不是String方法？因为需要保证每个err实例是不相同的（即使错误内容一样）
}
