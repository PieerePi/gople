// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type StringReader struct {
	str    string
	offset int64
}

func NewStringReader(s string) *StringReader {
	return &StringReader{s, 0}
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr == nil {
		return 0, fmt.Errorf("input StringReader is nil")
	}
	left := len(sr.str[sr.offset:])
	if left == 0 {
		return 0, io.EOF
	}
	// n = copy(p, sr.str[sr.offset:])
	if left > len(p) {
		n = len(p)
	} else {
		n = left
	}
	copy(p, sr.str[sr.offset:sr.offset+int64(n)])
	sr.offset += int64(n)
	return
}

// from go source code, src\strings\reader.go
func (sr *StringReader) Read2(b []byte) (n int, err error) {
	if sr.offset >= int64(len(sr.str)) {
		return 0, io.EOF
	}
	n = copy(b, sr.str[sr.offset:])
	sr.offset += int64(n)
	return
}

// !+
func main() {
	htmlText := `<html>
<head>
</head>
<body>
<div>
<img></img>
</div>
</body>
</html>`
	sr := NewStringReader(htmlText)
	doc, err := html.Parse(sr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

//!-
