// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func crawl(url string) []string {
	// fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type WORKLIST struct {
	list  []string
	depth int
}

type LINK struct {
	link  string
	depth int
}

var depthFlag = flag.Int("depth", 1, "the maximum depth of links")

// !+
func main() {
	flag.Parse()
	TOTALWORKERS := 20
	workListChan := make(chan WORKLIST)              // lists of URLs, may have duplicates
	unseenLinksChan := make(chan LINK, TOTALWORKERS) // de-duplicated URLs
	workDoneChan := make(chan int)
	working := 0

	// Add command-line arguments to worklist.
	go func() { workListChan <- WORKLIST{flag.Args(), 0}; workDoneChan <- 0 }()

	mainCtx, mainCancel := context.WithCancel(context.Background())

	// Create 20 crawler goroutines to fetch each unseen link.
	wg := &sync.WaitGroup{}
	for i := 0; i < TOTALWORKERS; i++ {
		wg.Add(1)
		go func(ctx context.Context, id int) {
			for {
				select {
				case link := <-unseenLinksChan:
					foundLinks := crawl(link.link)
					newdepth := link.depth + 1
					// go func() { worklist <- WORKLIST{foundLinks, newdepth} }()
					workListChan <- WORKLIST{foundLinks, newdepth}
					workDoneChan <- id
				case <-ctx.Done():
					// fmt.Printf("goroutine %d exits.\n", id)
					wg.Done()
					return
				}
			}
		}(mainCtx, i+1)
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	tick := time.Tick(1000 * time.Millisecond)
	var worklist []WORKLIST
	workerStatus := &strings.Builder{}
loop:
	for {
		select {
		case list := <-workListChan:
			if len(list.list) == 0 {
				continue
			}
			if list.depth >= *depthFlag {
				for _, link := range list.list {
					if !seen[link] {
						seen[link] = true
						fmt.Printf("%dth depth(the last): %s\n", list.depth, link)
					}
				}
				continue
			}
			worklist = append(worklist, list)
		case id := <-workDoneChan:
			if id != 0 {
				working--
			}
			if working == 0 && len(worklist) == 0 {
				mainCancel()
				wg.Wait()
				// fmt.Println("All crawl goroutins exit.")
				fmt.Println("All done.")
				fmt.Println("Worker status:")
				fmt.Print(workerStatus.String())
				break loop
			}
			if len(worklist) == 0 {
				continue
			}
			for i := 0; i < len(worklist) && working < TOTALWORKERS; i++ {
				ldepth := worklist[i].depth
				j := 0
				for ; j < len(worklist[i].list) && working < TOTALWORKERS; j++ {
					llink := worklist[i].list[j]
					if !seen[llink] {
						seen[llink] = true
						fmt.Printf("%dth depth: %s\n", ldepth, llink)
						working++
						unseenLinksChan <- LINK{llink, ldepth}
					}
				}
				if working == TOTALWORKERS {
					if j == len(worklist[i].list) {
						if i == len(worklist)-1 {
							worklist = nil
						} else {
							worklist = worklist[i+1:]
						}
					} else {
						worklist[i].list = worklist[i].list[j:]
						worklist = worklist[i:]
					}
					break
				}
			}
			if working < TOTALWORKERS {
				worklist = nil
			}
		case <-tick:
			fmt.Fprintf(workerStatus, "%v: %d workers are working.\n", time.Now(), working)
		}
	}
}

//!-
