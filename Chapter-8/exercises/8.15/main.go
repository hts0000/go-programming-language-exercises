package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// 练习 8.15
// 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞
// 修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好写
// 或者为每一个客户端的消息发出channel建立缓冲区，这样大部分的消息便不会被丢掉
// broadcaster应该用一个非阻塞的send向这个channel中发消息

func main() {
	// https://xingdl2007.gitbooks.io/gopl-soljutions/content/chapter-8-goroutines-and-channels.html
	// go broadcaster()
	// http.HandleFunc("/", hander)
	// log.Fatal(http.ListenAndServe("localhost:8001", nil))

	// https://github.com/ray-g/gopl/blob/master/ch08/ex8.14/chat.go
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	out chan<- string
	who string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-message:
			for cli := range clients {
				cli.out <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.out <- "Currently online:"
			for c := range clients {
				// exclude itself
				if c != cli {
					cli.out <- c.who
				}
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.out)
		}
	}
}

// https://xingdl2007.gitbooks.io/gopl-soljutions/content/chapter-8-goroutines-and-channels.html
func hander(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // 解析请求中的参数，将解析结果存入r.Form
		log.Print(err)
	}
	var who string
	// r.From 类型为 map[string][]string.参数可能有多个,通常为空格分隔,url中如何表示空格?
	// 用+或者直接用空格,但是要经过转换,golang中提供了转换函数url.QueryEscape()
	for k, v := range r.Form {
		if k == "name" { // 获取name参数
			who = v[0]
		}
	}
	if who == "" {
		who = r.RemoteAddr
	}
	hi, ok := w.(http.Hijacker) // 断言Hijacker,判断是否可以劫持
	if !ok {
		log.Fatalln("Can't Hijack")
	}
	conn, _, err := hi.Hijack() // 劫持链接，返回net.Conn,*bufio.ReadWriter,err,劫持者应该自行处理、关闭链接,*bufio.ReadWriter则可以对链接进行读写
	if err != nil {
		log.Fatalln("Hijack failed")
	}

	ch := make(chan string, 20) // 建立缓冲区
	go clientWriter(conn, ch)

	cli := client{ch, who}

	ch <- "You are " + who
	message <- who + "has arrived"
	entering <- cli

	// 一个小小的闭包，就能解决当前链接超时的问题
	timeout := time.NewTicker(5 * time.Minute) // 计时器
	go func() {
		<-timeout.C // 计时器到时间后会往这个管道塞进一个数据
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timeout.Reset(5 * time.Minute) // 重置计时器
		message <- who + ":" + input.Text()
	}

	leaving <- cli
	message <- who + "has left"
	conn.Close()
}

// https://github.com/ray-g/gopl/blob/master/ch08/ex8.14/chat.go
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	ch <- "Please type your name: "
	who := ""
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		who = scanner.Text()
	} else {
		ch <- "use default name: " + conn.RemoteAddr().String()
		who = conn.RemoteAddr().String()
	}

	cli := client{ch, who}

	ch <- "You are " + who
	message <- who + "has arrived"
	entering <- cli

	// 一个小小的闭包，就能解决当前链接超时的问题
	timeout := time.NewTicker(5 * time.Minute) // 计时器
	go func() {
		<-timeout.C // 计时器到时间后会往这个管道塞进一个数据
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timeout.Reset(5 * time.Minute) // 重置计时器
		message <- who + ":" + input.Text()
	}

	leaving <- cli
	message <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
