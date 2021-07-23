package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:10021")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go connection(conn)
	}
}

func connection(conn net.Conn) {
	defer func() {
		log.Printf("%v connection close\n", conn.LocalAddr())
		conn.Close()
	}()
	_, err := io.Copy(conn, conn)
	if err != nil {
		return
	}
}
