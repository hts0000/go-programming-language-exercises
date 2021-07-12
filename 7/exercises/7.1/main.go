package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// 练习 7.1
// 使用来自ByteCounter的思路，实现一个针对对单词和行数的计数器
// 你会发现bufio.ScanWords非常的有用

type ByteCounter int
type WordCounter int
type LineCounter int

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	str := "hello world! \nA 'Hello, World!' program generally is a computer program \nthat outputs or \ndisplays the message 'Hello, World!'."
	var w WordCounter
	w.Write([]byte(str))
	fmt.Println(w)
	w = 0
	fmt.Fprintf(&w, "'Hello, World!' program\n%s", str)
	fmt.Println(w)

	var l LineCounter
	l.Write([]byte(str))
	fmt.Println(l)
	l = 0
	fmt.Fprintf(&l, "'Hello, World!' program\n%s", str)
	fmt.Println(l)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("wordcounter: %v", err)
	}
	return int(*c), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("linecounter: %v", err)
	}
	return int(*c), nil
}
