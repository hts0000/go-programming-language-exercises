package main

import "io"

// 练习 7.5
// io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader
// 实现这个LimitReader函数：
// func LimitReader(r io.Reader, n int64) io.Reader

type LimitedReader struct {
	reader io.Reader
	limit  int64
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.limit {
		p = p[0:r.limit]
	}

	n, err = r.reader.Read(p)
	r.limit -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
