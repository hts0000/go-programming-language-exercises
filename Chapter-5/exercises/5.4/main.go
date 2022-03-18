package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 练习 5.4
// 扩展visit函数
// 使其能够处理其他类型的结点，如images、scripts和style sheets

// 未完成

func main() {
	doc, err := html.Parse(os.Stdin)
	// fmt.Println(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	visit(nil, doc)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
