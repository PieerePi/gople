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
go version go1.11 windows/amd64, 2.60GHz i5-3230M
go test -cpu=4 -bench=. gople\ch1\e1.3\strop_test.go -slen=10 -scount=1000
strop_test init done
goos: windows
goarch: amd64
BenchmarkStrOp1-4           1000           1503459 ns/op
BenchmarkStrOp2-4         100000             19074 ns/op
PASS
ok      command-line-arguments  4.301s
*/
