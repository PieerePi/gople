// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type DocStats struct {
	words int
	lines int
}

func (d *DocStats) Write(p []byte) (int, error) {
	lines := bufio.NewScanner(bytes.NewReader(p))
	// lines.Split(bufio.ScanLines)
	for lines.Scan() {
		d.lines++
	}
	words := bufio.NewScanner(bytes.NewReader(p))
	words.Split(bufio.ScanWords)
	for words.Scan() {
		d.words++
	}
	return len(p), nil
}

const descSplitFunc = `SplitFunc is the signature of the split function used to tokenize the input. The arguments are an initial substring of the remaining unprocessed data and a flag, atEOF, that reports whether the Reader has no more data to give. The return values are the number of bytes to advance the input and the next token to return to the user, if any, plus an error, if any.
Scanning stops if the function returns an error, in which case some of the input may be discarded.
Otherwise, the Scanner advances the input. If the token is not nil, the Scanner returns it to the user. If the token is nil, the Scanner reads more data and continues scanning; if there is no more data--if atEOF was true--the Scanner returns. If the data does not yet hold a complete token, for instance if it has no newline while scanning lines, a SplitFunc can return (0, nil, nil) to signal the Scanner to read more data into the slice and try again with a longer slice starting at the same point in the input.
The function is never called with an empty data slice unless atEOF is true. If atEOF is true, however, data may be non-empty and, as always, holds unprocessed text.`

func main() {
	//!+main
	var d DocStats
	d.Write([]byte(descSplitFunc))
	fmt.Println(d) // "{210 4}"

	d = DocStats{}
	var name = "GoLang"
	fmt.Fprintf(&d, "hello, %s\nhello, pieere\nafter a day of study\nbye, %[1]s, see you tomorrow\nsee you, pieere", name)
	fmt.Println(d) // "{17 5}"
	//!-main
}
