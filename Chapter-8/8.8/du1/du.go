package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var files int64 // 28240

func main() {
	var wg sync.WaitGroup
	fileSizes := make(chan int64)
	var totleSize int64
	wg.Add(1)
	go func() {
		for fileSize := range fileSizes {
			totleSize += fileSize
		}
		wg.Done()
	}()
	walkDir("/home/hts", fileSizes)
	close(fileSizes)
	wg.Wait()

	fmt.Println(totleSize)
	fmt.Printf("%.1f GB\t%v\n", float64(totleSize)/1e9, "/home/hts")
	fmt.Println(files)

	// flag.Parse()
	// roots := flag.Args()
	// if len(roots) == 0 {
	// 	roots = []string{"."}
	// }

	// fileSizes := make(chan int64)
	// go func() {
	// 	for _, root := range roots {
	// 		walkDir(root, fileSizes)
	// 	}
	// 	close(fileSizes)
	// }()

	// var nfiles, nbytes int64
	// for size := range fileSizes {
	// 	nfiles++
	// 	nbytes += size
	// }
	// printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d file  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		files++
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
