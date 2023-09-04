// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

// Package memotest provides common functions for
// testing various designs of the memo package.
package memo

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

// !+httpRequestBody
func httpGetBody(url string, index int, done <-chan struct{}) (interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if done != nil {
		go func() {
			for {
				select {
				case <-done:
					cancel()
					// fmt.Println("\t\t", index, url, "cancel exit")
					return
				case <-ctx.Done():
					// fmt.Println("\t\t", index, url, "normal exit")
					return
				}
			}
		}()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, index int, done <-chan struct{}) (interface{}, error)
	PrintKeys()
}

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

func Sequential(t *testing.T, m M) {
	//!+seq
	i := 0
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, i, nil)
		if err != nil {
			log.Print("\t\t RESULT: ", i, " ", url, " ", err)
			i++
			continue
		}
		fmt.Printf("\t\t RESULT: %d, %s, %s, %d bytes\n",
			i, url, time.Since(start), len(value.([]byte)))
		i++
	}
	fmt.Println("\t\t Cache keys:")
	m.PrintKeys()
	//!-seq
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

func Concurrent(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup
	i := 0
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string, index int) {
			defer n.Done()
			start := time.Now()
			done := make(chan struct{})
			go func() {
				time.Sleep(100 * time.Millisecond)
				if rand.Intn(8)%3 == 0 {
					close(done)
					fmt.Printf("\t\t Cancel %d, %s\n", index, url)
				}
			}()
			value, err := m.Get(url, index, done)
			if err != nil {
				log.Print("\t\t RESULT: ", index, " ", url, " ", err)
				return
			}
			fmt.Printf("\t\t RESULT: %d, %s, %s, %d bytes\n",
				index, url, time.Since(start), len(value.([]byte)))
		}(url, i)
		i++
	}
	n.Wait()
	fmt.Println("\t\t Cache keys:")
	m.PrintKeys()
	//!-conc
}
