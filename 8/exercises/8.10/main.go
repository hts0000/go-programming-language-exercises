package main

import (
	"fmt"
	"log"
	"os"

	"ex-8.10/links"
)

// 练习 8.10
// HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消
// 修改8.6节中的web crawler来支持取消http请求
// （提示：http.Get并没有提供方便地定制一个请求的方法。你可以用http.NewRequest来取而代之，设置它的Cancel字段，然后用http.DefaultClient.Do(req)来进行这个http请求。）

// 并发的退出
// 创建一个绝不会写入数据的channel done来广播通知所有goroutine是否需要结束
// 启动一个goroutine来检测是否收到退出信号，如果收到退出信号，则close(done)
// 编写一个cancall函数，select done是否能取到值
// 如果不行说明channel未关闭，返回false
// 如果可以取到说明channel已关闭（关闭的channel可以取到零值），返回true
// 在需要并发启动的函数中调用cancall函数，判断是否需要结束了
// 在需要并发的函数中调佣cancall是无侵入式的

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}
