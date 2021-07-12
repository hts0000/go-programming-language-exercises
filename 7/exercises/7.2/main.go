package main

import "io"

// 练习 7.2
// 写一个带有如下函数签名的函数CountingWriter
// 传入一个io.Writer接口类型
// 返回一个新的Writer类型把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类型指针
// func CountingWriter(w io.Writer) (io.Writer, *int64)

type CounterWriter struct {
	writer io.Writer
	count  int64
}

func main() {

}

func (w *CounterWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CounterWriter{
		writer: w,
		count:  0,
	}
	return cw, &cw.count
}
