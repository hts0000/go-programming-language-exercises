package main

import "fmt"

type IntList struct {
	Value int
	Tail  *IntList
}

func main() {
	// 表头
	head := &IntList{0, nil}

	// 初始化链表
	for i, h := 1, head; i < 11; i++ {
		h.Tail = NewNode(i)
		h = h.Tail
	}
	fmt.Println(head.Sum())
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func NewNode(x int) *IntList {
	return &IntList{x, nil}
}
