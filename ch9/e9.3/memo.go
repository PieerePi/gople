// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 278.

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

import (
	"context"
	"fmt"
	"strings"
)

//!+Func

// Func is the type of the function to memoize.
type Func func(key string, index int, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	index    int
	done     <-chan struct{}
	response chan<- result // the client wants a single result
}

type Memo struct {
	requests   chan request
	cache      map[string]*entry
	deleteChan chan string
}

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request), cache: make(map[string]*entry), deleteChan: make(chan string)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, index int, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, index, done, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) PrintKeys() {
	for key := range memo.cache {
		fmt.Println("\t\t", key)
	}
}

func (memo *Memo) Close() { close(memo.requests) }

//!-get

//!+monitor

func (memo *Memo) server(f Func) {
	for {
		select {
		case req, ok := <-memo.requests:
			if !ok {
				return
			}
			e := memo.cache[req.key]
			if e == nil {
				// This is the first request for this key.
				e = &entry{ready: make(chan struct{})}
				memo.cache[req.key] = e
				fmt.Println("\t\t", req.index, req.key, "invoke worker to get result")
				go e.call(f, req.key, req.index, req.done, memo.deleteChan) // call f(key)
			} else {
				fmt.Println("\t\t", req.index, req.key, "waiting result from cache")
			}
			go e.deliver(req.response)
		case key := <-memo.deleteChan:
			delete(memo.cache, key)
		}
	}
}

func (e *entry) call(f Func, key string, index int, done <-chan struct{}, deleteChan chan<- string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, index, done)
	select {
	case _, ok := <-done:
		// done and err, tell server not to cache it
		if !ok && strings.HasSuffix(e.res.err.Error(), context.Canceled.Error()) {
			go func() { deleteChan <- key }()
			fmt.Println("\t\t", index, key, "has been canceled, will not add to cache")
		}
	default:
	}
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}

//!-monitor
