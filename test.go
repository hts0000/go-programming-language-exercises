package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

// export EDITOR='program' 设置命令行默认编辑器
// export VISUAL='program' 设置GUI默认编辑器

func main() {
	// fileName := "test*.txt"

	// fp, err := os.CreateTemp("", fileName)
	// fmt.Println(fp.Name())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer fp.Close()
	// defer func() {
	// 	if os.Remove(fp.Name()).Error() != "" {
	// 		log.Printf("cant not remove tmp file %s, err: %v", fp.Name(), os.Remove(fp.Name()).Error())
	// 	}
	// }()
	// editorPath := `D:\Program Files\Sublime Text 3\sublime_text.exe`
	// editor := `sublime_text.exe`
	// cmd := &exec.Cmd{
	// 	Path:   editorPath,
	// 	Args:   []string{editor, fp.Name()},
	// 	Stdin:  os.Stdin,
	// 	Stdout: os.Stdout,
	// 	Stderr: os.Stderr,
	// }
	// // cmd := exec.Command(editorPath, fp.Name())
	// fmt.Println(cmd.Args)
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println()

	// cmd := &exec.Cmd{
	// 	Path:   "/usr/bin/vim",
	// 	Args:   []string{filepath.Base("/usr/bin/vim"), "/tmp/test.txt"},
	// 	Stdin:  os.Stdin,
	// 	Stdout: os.Stdout,
	// 	Stderr: os.Stderr,
	// }
	// fmt.Println(cmd.Args)
	// cmd.Run()
	// fmt.Println(cmd.Run().Error())

	// fmt.Println(4398046511618 & (1 << 0))
	// fmt.Printf("%#b\n", 4398046511618)
	// fmt.Printf("%#b\n", 4398046511616)
	// x := 4398046511618
	// x = x - ((x >> 1) & 0x5555555555555555)
	// x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	// x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	// x = x + (x >> 8)
	// x = x + (x >> 16)
	// x = x + (x >> 32)
	// fmt.Println(int(x & 0x7f))

	// n := -100.111
	// fmt.Println(+100.111)
	// fmt.Println(1 + n)
	// t2, _ := time.ParseDuration("-1h")
	// fmt.Println(time.Now().Add(t2))

	// tempdir, err := ioutil.TempDir("", "test")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(tempdir)
	// defer func() {
	// 	if err := os.RemoveAll(tempdir); err != nil {
	// 		log.Print(err)
	// 	}
	// }()

	// resp, err := http.Get("http://shouce.jb51.net/gopl-zh/ch8/ch8-06.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// doc, err := html.Parse(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// b := &bytes.Buffer{}
	// err = html.Render(b, doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file, err := os.Create("./index.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = io.Copy(file, b)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	links := crawlDepth("http://shouce.jb51.net/gopl-zh/ch8/ch8-06.html", 3)

	base, err := url.Parse("http://shouce.jb51.net/gopl-zh/ch8/ch8-06.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base.String(), base.Host, base.Path)

	for _, url := range linkURLs(links, base) {
		fmt.Println(url)
	}
}

// 练习 8.7
// 完成一个并发程序来创建一个线上网站的本地镜像，
// 把该站点的所有可达的页面都抓取到本地硬盘。
// 为了省事，我们这里可以只取出现在该域下的所有页面(比如golang.org结尾，译注：外链的应该就不算了。)
// 当然了，出现在页面里的链接你也需要进行一些处理,
// 使其能够在你的镜像站点上进行跳转，而不是指向原始的链接。

// 1.将传入url的网页，及其能到达的网页（深入深度3）下载到本地
// 		1.1 解析传入url为*html.Node，根据深度，爬取所有能到达的网页，解析为*html.Node，塞进a管道
// 		1.2 解析为baseurl，塞入b管道
// 		1.3 从b中取出baseurl，重写为本地可以访问的url
//

var urlChan = make(chan string)
var nodeChan = make(chan *html.Node)

func crawlDepth(baseurl string, depth uint8) []*html.Node {
	resp, err := http.Get(baseurl)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	links := linkNodes(doc)

	for _, link := range links {
		fmt.Printf("Attr = %#v, DataAtom = %v, Type = %v\n", link.Attr, link.DataAtom, link.Type)
	}

	return links
}

func linkNodes(doc *html.Node) []*html.Node {
	var links []*html.Node
	// 构建了一个闭包
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, n)
		}
	}
	forEachNode(doc, visitNode, nil)
	return links
}

// 从html解析树的头节点n开始，遍历这颗树获取所有节点，存放到links []*html.Node
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

func linkURLs(linkNodes []*html.Node, base *url.URL) []string {
	var urls []string
	for _, n := range linkNodes {
		for _, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			// ignore bad and non-local URLs
			if err != nil {
				log.Printf("skipping %q: %s", a.Val, err)
				continue
			}
			if link.Host != base.Host {
				//log.Printf("skipping %q: non-local host", a.Val)
				continue
			}
			urls = append(urls, link.String())
		}
	}
	return urls
}
