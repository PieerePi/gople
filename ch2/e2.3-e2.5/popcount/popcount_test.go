// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount

import (
	"testing"
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
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(0x1234567890ABCDEF)
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
go test -cpu=1 -bench .
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch2/e2.3-e2.5/popcount
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkPopCount               1000000000               0.6008 ns/op
BenchmarkPopCount2              139100662                8.808 ns/op
BenchmarkBitCount               1000000000               0.6040 ns/op
BenchmarkPopCountByShifting     27931594                42.54 ns/op
BenchmarkPopCountByClearing     39550442                27.76 ns/op
PASS
ok      github.com/PieerePi/gople/ch2/e2.3-e2.5/popcount        6.227s
go test -cpu=6 -bench .
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch2/e2.3-e2.5/popcount
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkPopCount-6                     1000000000               0.6313 ns/op
BenchmarkPopCount2-6                    129351421                8.349 ns/op
BenchmarkBitCount-6                     1000000000               0.5973 ns/op
BenchmarkPopCountByShifting-6           25940056                42.55 ns/op
BenchmarkPopCountByClearing-6           48317347                27.89 ns/op
PASS
ok      github.com/PieerePi/gople/ch2/e2.3-e2.5/popcount        6.309s
*/
