package main

import (
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
	// 启动一个goroutines，将命令行输入传输给连接
	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			log.Print(err)
		}
	}()
	// 从连接中读取数据，输出到命令行
	// if _, err := io.Copy(os.Stdout, conn); err != nil {
	// 	log.Print(err)
	// }
	for {
		res, err := io.ReadAll(conn)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(res)
	}

}
