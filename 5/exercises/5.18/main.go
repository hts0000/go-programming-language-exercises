package main

// 练习5.18
// 不修改fetch的行为，重写fetch函数，
// 要求使用defer机制关闭文件

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	filename, n, err := fetch("https://github.com/")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("download %s success, size %d byte\n", filename, n)
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	return local, n, err
}
