// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineInfo struct {
	count     int
	filenames map[string]bool
}

func main() {
	//var dupfiles []string
	counts := make(map[string]*lineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	dupfiles := make(map[string]bool) // only print file with duplicated lines once
	for _, lineinfo := range counts {
		if lineinfo.count > 1 {
			for filename := range lineinfo.filenames {
				if !dupfiles[filename] {
					fmt.Println(filename)
					dupfiles[filename] = true
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]*lineInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = &lineInfo{0, make(map[string]bool)}
		}
		counts[input.Text()].count++
		counts[input.Text()].filenames[f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
