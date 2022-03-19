package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// go run .\main.go http://www.baidu.com http://www.taobao.com
func main() {
	for i := 0; i < 2; i++ {
		start := time.Now()
		ch := make(chan string)
		for _, url := range os.Args[1:] {
			go fetch(url, ch)
		}

		for range os.Args[1:] {
			fmt.Println(<-ch)
		}
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
		time.Sleep(time.Second * 60)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
