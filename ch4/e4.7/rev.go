// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := [][]byte{
		[]byte("start"),
		[]byte("tert"),
		[]byte("12Ø34"),
		[]byte("123Ø4"),
		[]byte("1Ø234"),
		[]byte("¤2Ø34"),
		[]byte("1¤3Ø4"),
		[]byte("1Ø2¤4"),
		[]byte("Øe¤¥næn"),
		[]byte("Hello, 世界"),
	}

	for i := range data {
		fmt.Printf("%s | ", data[i])
		ReverseUtf8(data[i])
		fmt.Printf("%s | ", data[i])
		ReverseUtf8(data[i])
		fmt.Printf("%s\n", data[i])
	}
}

//!+rev

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func rotateLeftBytes(b []byte, p int) {
	reverseBytes(b[:p])
	reverseBytes(b[p:])
	reverseBytes(b)
}

// This version of RotateFtBytes requires extra space, but results in poorer performance.
// See rotate.go and rotate_test.go for more information.
// func rotateLeftBytes(b []byte, p int) {
// 	blen := len(b)
// 	if p <= 0 || p >= blen {
// 		return
// 	}
// 	left := make([]byte, p)
// 	copy(left, b[:p])
// 	copy(b, b[p:])
// 	copy(b[blen-p:], left)
// }

// ReverseUtf8 reverses a UTF-8 byte slice in place.
func ReverseUtf8(b []byte) {
	if len(b) == 0 {
		return
	}

	r, size := utf8.DecodeRune(b)

	if r == utf8.RuneError {
		panic("Rune Error")
	}

	rotateLeftBytes(b, size)
	ReverseUtf8(b[:len(b)-size])
}

//!-rev
