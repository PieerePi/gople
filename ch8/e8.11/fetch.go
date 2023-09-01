// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 148.

// Fetch saves the contents of a URL into a local file.
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
)

// !+
// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(ctx context.Context, url string) (filename string, n int64, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

//!-

// go run fetch.go http://www.gopl.io/ http://www.gopl.io/reviews.html http://www.gopl.io/ch1.pdf
func main() {
	responses := make(chan string, len(os.Args[1:]))
	mainCtx, mainCancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	for _, url := range os.Args[1:] {
		wg.Add(1)
		lurl := url
		go func(ctx context.Context) {
			local, n, err := fetch(ctx, lurl)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch %s: %v\n", lurl, err)
			} else {
				responses <- lurl
				fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", lurl, local, n)
			}
			wg.Done()
		}(mainCtx)
	}
	<-responses
	mainCancel()
	wg.Wait()
}
