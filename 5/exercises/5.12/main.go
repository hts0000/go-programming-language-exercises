package main

import "golang.org/x/net/html"

// 练习5.12
// gopl.io/ch5/outline2（5.5节）的startElement和endElement共用了全局变量depth
// 将它们修改为匿名函数，使其共享outline中的局部变量

func main() {

}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
