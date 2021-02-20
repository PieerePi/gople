// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"github.com/PieerePi/gople/ch2/e2.3-e2.5/popcount"
)

// -- Alternative implementations --

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount2(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

/*
go version go1.15.7 windows/amd64, 2.1-4GHz R5-4600U
go test -cpu=1 -bench=. popcount_test.go
goos: windows
goarch: amd64
BenchmarkPopCount               1000000000               0.272 ns/op
BenchmarkPopCount2              153538957                7.86 ns/op
BenchmarkBitCount               1000000000               0.259 ns/op
BenchmarkPopCountByShifting     33558296                35.7 ns/op
BenchmarkPopCountByClearing     92298462                13.2 ns/op
PASS
ok      command-line-arguments  6.168s
go test -cpu=6 -bench=. popcount_test.go
goos: windows
goarch: amd64
BenchmarkPopCount-6                     1000000000               0.277 ns/op
BenchmarkPopCount2-6                    146317784                8.24 ns/op
BenchmarkBitCount-6                     1000000000               0.257 ns/op
BenchmarkPopCountByShifting-6           31427022                36.8 ns/op
BenchmarkPopCountByClearing-6           92288523                13.4 ns/op
PASS
ok      command-line-arguments  5.340s
*/
