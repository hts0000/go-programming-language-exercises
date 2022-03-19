package main

import (
	"bufio"
	"os"
	"testing"
)

// go test --bench=. --benchmem
// cpu: Intel(R) Core(TM) i5-5200U CPU @ 2.20GHz
// BenchmarkEcho1-4             151           7985130 ns/op        24546233 B/op       2879 allocs/op
// BenchmarkEcho2-4             141           8155449 ns/op        24546234 B/op       2879 allocs/op
// BenchmarkEcho3-4           18657             60840 ns/op           16384 B/op          1 allocs/op
// PASS
// ok      test/Chapter-1/exercises/1.3    6.151s

var data []string

func init() {
	data = generateData()
}

func generateData() (res []string) {
	filename := `.\Robert Louis Stevenson.txt`
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			panic(os.ErrNotExist)
		}
		panic(err)
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(data)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(data)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(data)
	}
}
