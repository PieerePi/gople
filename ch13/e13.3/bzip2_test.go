// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bzip

import (
	"bufio"
	"bytes"
	"compress/bzip2" // reader
	"fmt"
	"io"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := NewWriter(&compressed)

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	// truncate --size 0 hello.txt;
	// for ((i=0;i<1000000;i++)) do echo -n "hello" >> hello.txt; done;
	// bzip2 -f hello.txt; ls -l hello.txt.bz2 | awk '{print $5}'
	// Terrible!!!
	// bzip2 compress 5000000 bytes (1000000 times "hello") to 255 bytes
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}

func TestConcurrentWrites(t *testing.T) {
	// In a container, we should use the following method.
	// import _ "github.com/uber-go/automaxprocs"
	// WORKERS := runtime.GOMAXPROCS(0)
	WORKERS := runtime.NumCPU()
	n := 1000
	c := make(chan int, n)
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)

	compressed := &bytes.Buffer{}
	w := NewWriter(compressed)
	var errs = make([]error, WORKERS)
	wg := &sync.WaitGroup{}
	// Use buffer instead of fmt to record the worker details
	details := &bytes.Buffer{}

	consume := func(ie int) {
		defer wg.Done()
		for i := range c {
			details.Write([]byte(fmt.Sprintf("worker-#%d - %d\n", ie, i)))
			_, errs[ie] = w.Write([]byte(fmt.Sprintf("%d\n", i)))
			if errs[ie] != nil {
				return
			}
		}
	}
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go consume(i)
	}
	wg.Wait()
	for i := 0; i < WORKERS; i++ {
		if errs[i] != nil {
			t.Errorf("%s", errs[i])
		}
	}
	w.Close()

	// Worker details
	fmt.Println(details.String())

	// Check each number is present.
	seen := make(map[int]bool)
	decompressed := &bytes.Buffer{}
	io.Copy(decompressed, bzip2.NewReader(compressed))
	s := bufio.NewScanner(decompressed)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			t.Errorf("%s", err)
			return // Corrupted writes?
		}
		seen[i] = true
	}
	missing := &bytes.Buffer{}
	for i := 0; i < n; i++ {
		if !seen[i] {
			missing.Write([]byte(fmt.Sprintf("%d\n", i)))
		}
	}
	if missing.Len() > 0 {
		t.Errorf("missing: \n%s", missing.String())
	}
}
