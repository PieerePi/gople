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
without sync.Once

go test -cpu=1 -bench .
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch9/e9.2
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkPopCount               1000000000               0.5908 ns/op
BenchmarkPopCount2              147000634                8.310 ns/op
BenchmarkBitCount               1000000000               0.5904 ns/op
BenchmarkPopCountByShifting     27595278                42.03 ns/op
BenchmarkPopCountByClearing     39042673                27.46 ns/op
PASS
ok      github.com/PieerePi/gople/ch9/e9.2      6.078s

go test -cpu=6 -bench .
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch9/e9.2
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkPopCount-6                     1000000000               0.6259 ns/op
BenchmarkPopCount2-6                    126873954                8.418 ns/op
BenchmarkBitCount-6                     1000000000               0.6062 ns/op
BenchmarkPopCountByShifting-6           27262131                43.10 ns/op
BenchmarkPopCountByClearing-6           44748902                27.51 ns/op
PASS
ok      github.com/PieerePi/gople/ch9/e9.2      6.255s

------

with sync.Once

go test -cpu=1 -bench .
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch9/e9.2
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkPopCount               174582782                6.870 ns/op
BenchmarkPopCount2              57802353                19.05 ns/op
BenchmarkBitCount               1000000000               0.5979 ns/op
BenchmarkPopCountByShifting     26439812                43.63 ns/op
BenchmarkPopCountByClearing     28444713                44.43 ns/op
PASS
ok      github.com/PieerePi/gople/ch9/e9.2      6.568s

go test -cpu=6 -bench .
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch9/e9.2
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkPopCount-6                     165485512                6.896 ns/op
BenchmarkPopCount2-6                    62817357                19.86 ns/op
BenchmarkBitCount-6                     1000000000               0.5988 ns/op
BenchmarkPopCountByShifting-6           27516687                42.47 ns/op
BenchmarkPopCountByClearing-6           26852305                41.71 ns/op
PASS
ok      github.com/PieerePi/gople/ch9/e9.2      6.591s
*/
