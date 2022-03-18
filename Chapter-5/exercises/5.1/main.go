package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 练习 5.1
// 修改findlinks代码中遍历n.FirstChild链表的部分
// 将循环调用visit，改成递归调用

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
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				// links = append(links, a.Val)
				fmt.Println(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
