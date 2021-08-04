package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

// 编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来

var (
	sema = make(chan struct{}, 20)
)

func main() {
	rootPath := "/root"
	entries, err := ioutil.ReadDir(rootPath)
	if os.IsPermission(err) { // "permission denied"
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return
	}
	dirs := make([]string, 0)
	for _, entrie := range entries {
		if entrie.IsDir() {
			dirs = append(dirs, filepath.Join(rootPath, entrie.Name()))
		}
	}
	sort.Strings(dirs)
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}

func dirSize(dir string, dirsSize map[string]int64) {
	sema <- struct{}{}
	defer func() { <-sema }()
	dirsSize[dir] = size(dir)
}

func size(dir string) int64 {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			size(subdir)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dirents: %v\n", err)
		return nil
	}
	return entries
}

// 1. 获取root下所有目录（权限判断），存储在[]string里（每隔一段时间重新获取），并排序
// 2. 为map的所有key分别创建goroutine（用信号量限制20个），递归计算这个目录的大小，将大小更新到map value中
// 3. 启动一个goroutine等待所有目录统计完成
