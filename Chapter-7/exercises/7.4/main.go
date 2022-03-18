package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 练习 7.4
// strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）
// 实现一个简单版本的NewReader，并用它来构造一个接收字符串输入的HTML解析器（§5.2）

type StringReader struct {
	s string
}

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}

	r := NewStringReader(os.Args[1])
	p := make([]byte, 100)
	r.Read(p)
	fmt.Println("##", string(p))
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, []byte(r.s)) // string类型可以等价为[]byte类型
	fmt.Println(n)
	if len(r.s[n:]) == 0 {
		err = io.EOF
	}
	return
}

func NewStringReader(s string) *StringReader {
	return &StringReader{s}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %V", url, err)
	}
	return visit(nil, doc), nil
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
