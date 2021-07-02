package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 练习 5.3
// 编写函数输出所有text结点的内容
// 注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。

func main() {
	doc, err := html.Parse(os.Stdin)
	// fmt.Println(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	// for _, link := range visit(nil, doc) {
	// 	fmt.Println(link)
	// }
	visit(nil, doc)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
