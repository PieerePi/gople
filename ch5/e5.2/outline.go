// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	elements := make(map[string]int)
	outline(elements, doc)
	for tag, count := range elements {
		fmt.Printf("%q appears %d times.\n", tag, count)
	}
}

func outline(elem map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elem[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(elem, c)
	}
}

//!-
