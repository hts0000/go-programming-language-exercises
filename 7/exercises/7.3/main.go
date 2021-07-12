package main

import (
	"bytes"
	"fmt"
)

//练习 7.3
// 为在gopl.io/ch4/treesort (§4.4)的*tree类型
// 实现一个String方法去展示tree类型的值序列

type tree struct {
	value       int
	left, right *tree
}

func main() {
	values := []int{3, 5, 6, 8, 1, 2, 19, 77, 4, 18}
	root := Sort(values)
	fmt.Println(values)

	fmt.Println(root) // 隐式调用tree类型的String方法
}

func (t *tree) String() string {
	order := appendValues([]int{}, t)
	if len(order) == 0 {
		return "[]"
	}

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v", order)
	return b.String()
}

func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	// fmt.Println(root.value)
	// 传入一个底层数组与values相同的新切片
	appendValues(values[:0], root)
	// fmt.Printf("v1 = %p, v2 = %p, v1cap = %d, v2cap = %d\n", values, values[:0], cap(values), cap(values[:0]))
	// fmt.Printf("v1len = %d, v2len = %d\n", len(values), len(values[:0]))
	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		// 按照前序顺序将生成的二叉树的值append到新切片中
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 生成一颗排序二叉树
// 根节点总是第一个传入的元素
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
