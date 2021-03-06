// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

package main

import "fmt"

//!+
func maxWithError(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("max: too few arguments")
	}
	m := vals[0]
	for _, val := range vals[1:] {
		if val > m {
			m = val
		}
	}
	return m, nil
}

func minWithError(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("min: too few arguments")
	}
	m := vals[0]
	for _, val := range vals[1:] {
		if val < m {
			m = val
		}
	}
	return m, nil
}

func max(m int, vals ...int) int {
	for _, val := range vals {
		if val > m {
			m = val
		}
	}
	return m
}

func min(m int, vals ...int) int {
	for _, val := range vals {
		if val < m {
			m = val
		}
	}
	return m
}

//!-

func main() {
	//!+main
	fmt.Println(maxWithError())  //  "0 max: too few arguments"
	fmt.Println(minWithError())  //  "0 min: too few arguments"
	fmt.Println(max(1, 2, 3, 4)) //  "4"
	fmt.Println(min(1, 2, 3, 4)) //  "1"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(max(30, values...)) // "30"
	fmt.Println(min(-3, values...)) // "-3"
	//!-slice
}
