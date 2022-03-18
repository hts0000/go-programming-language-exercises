package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if ok := strings.HasPrefix(url, "http://"); !ok {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "execirses 1.9: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(url, resp.Status)
	}
}
