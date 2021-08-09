package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 练习 8.12
// 使broadcaster能够将arrival事件通知当前所有的客户端。
// 为了达成这个目的，你需要有一个客户端的集合，并且在entering和leaving的channel中记录客户端的名字

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

	input := bufio.NewScanner(conn)
	for input.Scan() {
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
