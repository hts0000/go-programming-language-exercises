package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// 练习 8.13
// 使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端
// 提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像input.Scanner()所做的那样

func main() {
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

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
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
