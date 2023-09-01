// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

// !+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	var fileSizes []chan int64
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		fileSizes = append(fileSizes, make(chan int64))
		go walkDir(root, &n, fileSizes[i])
	}
	//!-

	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
	var m sync.WaitGroup
	for i := range roots {
		m.Add(1)
		go func(mp *sync.WaitGroup, i int) {
			// Print the results periodically.
			var tick <-chan time.Time
			if *vFlag {
				tick = time.Tick(500 * time.Millisecond)
			}
		loop:
			for {
				select {
				case size, ok := <-fileSizes[i]:
					if !ok {
						break loop // fileSizes was closed
					}
					nfiles[i]++
					nbytes[i] += size
				case <-tick:
					printDiskUsage(roots[i], nfiles[i], nbytes[i])
				}
			}
			printDiskUsage(roots[i], nfiles[i], nbytes[i])
			mp.Done()
		}(&m, i)
	}

	n.Wait()
	for i := range roots {
		close(fileSizes[i])
	}

	m.Wait()

	printAllDiskUsage(nfiles, nbytes) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files  %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}

func printAllDiskUsage(nfiles, nbytes []int64) {
	var allfiles, allbytes int64
	for i := range nfiles {
		allfiles += nfiles[i]
	}
	for i := range nbytes {
		allbytes += nbytes[i]
	}
	fmt.Printf("All: %d files  %.1f GB\n", allfiles, float64(allbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
// !+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			if fi, err := entry.Info(); err == nil {
				fileSizes <- fi.Size()
			}
		}
	}
}

//!-walkDir

// !+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []fs.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
