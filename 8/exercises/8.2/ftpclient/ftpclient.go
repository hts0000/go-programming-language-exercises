package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:10021")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 从连接中读取数据，输出到命令行
	// go func() {
	// 	if _, err := io.Copy(os.Stdout, conn); err != nil {
	// 		log.Print(err)
	// 	}
	// }()

	readInput := bufio.NewReader(os.Stdin)
	readConn := bufio.NewReader(conn)
	for {
		fmt.Fprint(os.Stdout, "> ")
		input, err := readInput.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue  
		}
		_, err = io.WriteString(conn, input)
		if err != nil {
			log.Print(err)
			continue
		}
		r := make([]byte, 0)
		_, err = readConn.Read(r)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Fprint(os.Stdout, string(r))
		// if _, err := io.Copy(os.Stdout, conn); err != nil {
		// 	log.Print(err)
		// }
	}
}
