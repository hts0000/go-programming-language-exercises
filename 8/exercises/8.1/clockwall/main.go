package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var (
		ny, ld, sh string
	)
	flag.StringVar(&sh, "Shanghai", "localhost:8000", "set Shanghai timezone server")
	flag.StringVar(&ny, "NewYork", "localhost:8001", "set NewYork timezone server")
	flag.StringVar(&ld, "London", "localhost:8002", "set London timezone server")
	flag.Parse()

	for _, zone := range []string{sh, ny, ld} {
		conn, err := net.Dial("tcp", zone)
		if err != nil {
			log.Fatal(err)
		}
		go printZone(conn)
	}
	wg.Wait()
}

func printZone(conn net.Conn) {
	wg.Add(1)
	defer conn.Close()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Print(err)
	}
	wg.Done()
}
