package bank3

import "sync"

var (
	// 惯例来说，被mutex所保护的变量是在mutex变量声明之后立刻声明的
	// 如果你的做法和惯例不符，确保在文档里对你的做法进行说明。
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
