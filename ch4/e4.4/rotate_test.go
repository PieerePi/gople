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
go version go1.15.7 windows/amd64, 2.60GHz i5-4210M
go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch4/e4.4
BenchmarkRotateLeft-4           50032938                20.5 ns/op
BenchmarkRotateLeft2-4          23002650                46.8 ns/op
PASS
ok      github.com/PieerePi/gople/ch4/e4.4      2.498s
*/
