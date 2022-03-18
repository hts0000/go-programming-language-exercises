package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	ln := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLine(os.Stdin, counts, "os.Stdin", ln)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLine(f, counts, file, ln)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s:\t%d\t%s\n", ln[line], n, line)
		}
	}
}

func countLine(f *os.File, counts map[string]int, filename string, ln map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		ln[input.Text()] = filename
	}
}
