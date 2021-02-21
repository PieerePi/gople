// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaFloat(os.Args[i]))
	}
}

func commaFloat(s string) string {
	var sign string
	var fraction string
	if len(s) != 0 {
		if s[0] == '+' || s[0] == '-' {
			sign = s[0:1]
			s = s[1:]
		}
		if point := strings.LastIndex(s, "."); point != -1 {
			fraction = s[point:]
			s = s[0:point]
		}
	}
	return sign + comma(s) + fraction
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	i, j := n%3, n/3
	if i != 0 {
		buf.WriteString(s[:i] + ",")
	}
	for k := 0; k < j; k++ {
		buf.WriteString(s[k*3+i : k*3+i+3])
		if k != j-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

//!-
