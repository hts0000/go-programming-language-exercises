package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 练习5.17
// 编写多参数版本的ElementsByTagName
// 函数接收一个HTML结点树以及任意数量的标签名
// 返回与这些标签名匹配的所有元素
// 下面给出了2个例子：
// func ElementsByTagName(doc *html.Node, name...string) []*html.Node
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

func main() {
	url := "https://juejin.cn/post/6844904166586892295"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "5.17: %v\n", err)
		os.Exit(1)
	}
	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	for _, image := range images {
		fmt.Println(image)
	}

	for _, heading := range headings {
		fmt.Println(heading)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if len(name) == 0 {
		return nil
	}
	if doc.Type == html.ElementNode {
		for _, tag := range name {
			if doc.Data == tag {
				nodes = append(nodes, doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, name...)...)
	}
	return nodes
}
