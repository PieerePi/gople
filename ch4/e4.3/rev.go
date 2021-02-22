// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array
}

//!+rev

const (
	arraySize = 6
)

// reverse reverses an array of ints in place.
func reverse(pa *[arraySize]int) {
	for i, j := 0, len(pa)-1; i < j; i, j = i+1, j-1 {
		pa[i], pa[j] = pa[j], pa[i]
	}
}

//!-rev
