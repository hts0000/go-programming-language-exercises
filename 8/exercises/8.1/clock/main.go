package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var (
		port     string
		timeZone string
	)
	flag.StringVar(&port, "p", "8000", "set clock port")

	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, timeZone)
	}
}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	for {
		zone, _ := time.Now().Zone()
		_, err := io.WriteString(c, zone+": "+time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
