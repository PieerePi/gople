// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y, z IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Elems8()) // "[1 9 42 144]"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	x.Intersection(&y)
	fmt.Println(x.String()) // "{9 42}"

	x.Remove(9) // "{42}"
	x.Add(22)   // "{22, 42}"
	x.Difference(&y)
	fmt.Println(x.String()) // "{22}"

	x.Intersection(&z)
	fmt.Println(x.String()) // "{}"

	x.AddAll(9, 22) // "{9, 22}"
	x.SymmetricDifference(&y)
	fmt.Println(x.String()) // "{22, 42}"

	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// [1 9 42 144]
	// true false
	// {9 42}
	// {22}
	// {}
	// {22 42}

}

func Example_two() {
	var x IntSet
	x.Add(0)
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{0 1 9 42 144}"
	fmt.Println(x.String()) // "{0 1 9 42 144}"
	fmt.Println(x)          // "{[4398046511619 0 65536]}"
	fmt.Println(x.Has(0))   // "true"
	fmt.Println(x.Len())    // "5"

	x.Remove(9)
	fmt.Println(x.String()) // "{0 1 42 144}"
	fmt.Println(x.Len())    // "4"
	x.Add(9)
	fmt.Println(x.String()) // "{0 1 9 42 144}"
	fmt.Println(x.Len())    // "5"

	y := x.Copy()
	fmt.Println(y.String()) // "{0 1 9 42 144}"
	fmt.Println(y.Len())    // "5"
	x.Clear()
	fmt.Println(x.String()) // "{}"
	fmt.Println(x.Len())    // "0"
	z := x.Copy()
	fmt.Println(z.String()) // "{}"
	fmt.Println(z.Len())    // "0"
	fmt.Println(y.String()) // "{0 1 9 42 144}"
	fmt.Println(y.Len())    // "5"
	fmt.Println(y.Elems())  // "[1 9 42 144]"

	x.Add(100)
	fmt.Println(x.String()) // "{100}"
	fmt.Println(x.Len())    // "1"

	//!-note

	// Output:
	// {0 1 9 42 144}
	// {0 1 9 42 144}
	// {[4398046511619 0 65536]}
	// true
	// 5
	// {0 1 42 144}
	// 4
	// {0 1 9 42 144}
	// 5
	// {0 1 9 42 144}
	// 5
	// {}
	// 0
	// {}
	// 0
	// {0 1 9 42 144}
	// 5
	// [0 1 9 42 144]
	// {100}
	// 1

}

func Benchmark_Elems(b *testing.B) {
	var x IntSet
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		x.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = x.Elems()
	}
}

func Benchmark_Elems2(b *testing.B) {
	var x IntSet
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		x.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = x.Elems2()
	}
}

func Benchmark_Elems3(b *testing.B) {
	var x IntSet
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		x.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = x.Elems3()
	}
}

func Benchmark_Elems4(b *testing.B) {
	var x IntSet
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		x.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = x.Elems4()
	}
}

func Benchmark_Elems5(b *testing.B) {
	var x IntSet
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		x.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = x.Elems5()
	}
}

func Benchmark_Elems6(b *testing.B) {
	var x IntSet
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		x.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = x.Elems6()
	}
}

func Benchmark_Elems7(b *testing.B) {
	var x IntSet
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		x.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = x.Elems7()
	}
}

func Benchmark_Elems8(b *testing.B) {
	var x IntSet
	b.StopTimer()
	for i := 0; i < 10000; i++ {
		x.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = x.Elems8()
	}
}

/*
以下性能总结：
8>6>7，5>4，循环尽量使用range（性能好些但差别很小），用Go的方式写代码
2>3，用slice索引而不使用append，性能提升很小，也尽量用append，用Go的方式写代码
真正的差距在于slice空间的预分配
切记，尽量使用range和append，不要用c的方式写Go代码
*/
/*
go test -bench .
goos: windows
goarch: amd64
pkg: github.com/PieerePi/gople/ch6/e6.5
cpu: AMD Ryzen 5 4600U with Radeon Graphics
Benchmark_Elems-12         18459             66591 ns/op
Benchmark_Elems2-12        18774             62666 ns/op
Benchmark_Elems3-12        18142             66406 ns/op
Benchmark_Elems4-12        20691             59486 ns/op
Benchmark_Elems5-12        18120             67316 ns/op
Benchmark_Elems6-12         8145            145626 ns/op
Benchmark_Elems7-12         8083            153779 ns/op
Benchmark_Elems8-12         8737            145049 ns/op
PASS
ok      github.com/PieerePi/gople/ch6/e6.5      13.474s
*/
