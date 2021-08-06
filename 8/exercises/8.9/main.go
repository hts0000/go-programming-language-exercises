package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

// 练习 8.9
// 编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来

var wg sync.WaitGroup
var mu sync.RWMutex

func main() {
	for {
		dirs := make([]string, 0)
		for _, entrie := range dirents("/root") {
			if entrie.IsDir() {
				dirs = append(dirs, filepath.Join("/root", entrie.Name()))
			}
		}
		sort.Strings(dirs)

		dirsSize := make(map[string]int64)
		sema := make(chan struct{}, 3)
		for _, dir := range dirs {
			sema <- struct{}{}
			wg.Add(1)
			go size(dir, dirsSize, sema)
		}

		wg.Wait()

		for _, dir := range dirs {
			fmt.Printf("%s    %.1fM\n", dir, float64(dirsSize[dir])/1e6)
			// fmt.Printf("%s    %v\n", dir, dirsSize[dir])
		}
		time.Sleep(1 * time.Second)
	}
}

func size(dir string, dirsSize map[string]int64, sema chan struct{}) {
	size := walkDir(dir)
	mu.Lock()
	dirsSize[dir] = size
	mu.Unlock()
	<-sema
	wg.Done()
}

func walkDir(dir string) (size int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			size += walkDir(subdir)
		} else {
			size += entry.Size()
		}
	}
	return size
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if os.IsPermission(err) { // "permission denied"
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

// 1. 获取root下所有目录（权限判断），存储在[]string里（每隔一段时间重新获取），并排序
// 2. 为map的所有key分别创建goroutine（用信号量限制20个），递归计算这个目录的大小，将大小更新到map value中
// 3. 启动一个goroutine等待所有目录统计完成
