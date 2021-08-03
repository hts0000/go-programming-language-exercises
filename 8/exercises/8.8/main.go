package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	inputCh := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			inputCh <- input.Text()
		}
	}()

	for {
		select {
		case input := <-inputCh:
			go echo(c, input, 1*time.Second)
		case <-time.After(10 * time.Second):
			log.Print("no response in 10s, close client:", c.RemoteAddr())
			c.Close()
			return
		}
	}
}
