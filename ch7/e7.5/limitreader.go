package main

import "io"

type MyLimitReader struct {
	r io.Reader
	n int64
}

func NewMyLimitReader(r io.Reader, n int64) io.Reader {
	return &MyLimitReader{r, n}
}

func (lr *MyLimitReader) Read(p []byte) (n int, err error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	if lr.n < int64(len(p)) {
		n, err = lr.r.Read(p[:lr.n])
	} else {
		n, err = lr.r.Read(p)
	}
	lr.n -= int64(n)
	return
}

// from go source code, src\io\io.go
func (lr *MyLimitReader) Read2(p []byte) (n int, err error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > lr.n {
		p = p[0:lr.n]
	}
	n, err = lr.r.Read(p)
	lr.n -= int64(n)
	return
}
