package strop_test

import (
	"flag"
	"fmt"
	"strings"
	"testing"
)

// for BenchmarkStrOp1, the higher the value, the worse the performance
// try 100 1000 10000
var scount = flag.Int("scount", 1000, "total string count")

// try 1 10 100
var slen = flag.Int("slen", 10, "single string length")

var longStrArray []string

func init() {
	var str = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	// flag.Parse will fail if doesn't Init package testing
	testing.Init()
	flag.Parse()
	longStrArray = make([]string, *scount)
	for i := range longStrArray {
		for j := 0; j < *slen; j++ {
			longStrArray[i] += str[i%10]
		}
	}
	fmt.Println("strop_test init done")
}

// BenchmarkStrOp1 for test
func BenchmarkStrOp1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var result string // result should be declared here, in the loop
		for _, v := range longStrArray {
			result += v + " "
		}
	}
}

// BenchmarkStrOp2 join test
func BenchmarkStrOp2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(longStrArray, " ")
	}
}

/*
go test -cpu=1 -bench . -slen=10 -scount=1000
strop_test init done
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch1/e1.3
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkStrOp1              679           1674839 ns/op
BenchmarkStrOp2            62929             19239 ns/op
PASS
ok      github.com/PieerePi/gople/ch1/e1.3      3.117s
go test -cpu=6 -bench . -slen=10 -scount=1000
strop_test init done
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch1/e1.3
cpu: AMD Ryzen 5 4600U with Radeon Graphics
BenchmarkStrOp1-6            553           2187795 ns/op
BenchmarkStrOp2-6          58640             20420 ns/op
PASS
ok      github.com/PieerePi/gople/ch1/e1.3      3.213s
*/
