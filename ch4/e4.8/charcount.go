// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type isItem struct {
	isName string
	isFunc func(rune) bool
}

var isItmes = []isItem{
	{"space", unicode.IsSpace},
	{"symbol", unicode.IsSymbol},
	{"mark", unicode.IsMark},
	{"digit", unicode.IsDigit},
	{"print", unicode.IsPrint},
	{"punct", unicode.IsPunct},
	{"letter", unicode.IsLetter},
	{"number", unicode.IsNumber},
	{"control", unicode.IsControl},
	{"graphic", unicode.IsGraphic},
}

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	cats := make(map[string]int)
	cats2 := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		// ^Z on Windows
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for catName, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				cats[catName]++
			}
		}
		for _, item := range isItmes {
			if item.isFunc(r) {
				cats2[item.isName]++
			}
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Printf("\n%-30s count\n", "category")
	for cat, n := range cats {
		fmt.Printf("%-30s %d\n", cat, n)
	}
	fmt.Printf("\n%-30s count2\n", "category2")
	for cat, n := range cats2 {
		fmt.Printf("%-30s %d\n", cat, n)
	}
}

//!-
