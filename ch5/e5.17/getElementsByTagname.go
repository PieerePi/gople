// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: getElementsByTagName url names")
		os.Exit(1)
	}

	if err := getElementsByTagName(os.Args[1], os.Args[2:]); err != nil {
		fmt.Printf("getElementsByTagName failed: %v\n", err)
		os.Exit(1)
	}
}

func getElementsByTagName(url string, name []string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Get Failed: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	node := ElementsByTagName(doc, name...)
	if node == nil {
		fmt.Printf("No element with names given %v found\n", name)
	} else {
		fmt.Printf("Found %d elements with names given %v\n", len(node), name)
	}
	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Printf("found %d images, %d headings", len(images), len(headings))

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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

//!-forEachNode

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var node []*html.Node
	var nameMap = make(map[string]bool, len(name))

	for _, n := range name {
		nameMap[n] = true
	}

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && nameMap[n.Data] {
			node = append(node, n)
		}
	}, nil)

	return node
}
