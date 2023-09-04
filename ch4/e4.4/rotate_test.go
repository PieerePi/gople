package main

import (
	"testing"
)

func BenchmarkRotateLeft(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		rotateLeft(s, i%10)
	}
}

func BenchmarkRotateLeft2(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		rotateLeft2(s, i%10)
	}
}

/*
go test -bench .
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch4/e4.4
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkRotateLeft-12          48922082                21.49 ns/op
BenchmarkRotateLeft2-12         14273750                79.40 ns/op
PASS
ok      github.com/PieerePi/gople/ch4/e4.4      2.714s
*/
