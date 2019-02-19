// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

var tagAttrFilter = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
	"link":   "href", // should check rel or type attribute
}

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	for tag, attr := range tagAttrFilter {
		if n.Type == html.ElementNode && n.Data == tag {
			if n.Data == "link" {
				isStyleSheet := false
				link := ""
				for _, a := range n.Attr {
					if a.Key == "rel" && a.Val == "stylesheet" {
						isStyleSheet = true
					} else if a.Key == "type" && a.Val == "text/css" {
						isStyleSheet = true
					} else if a.Key == attr {
						link = a.Val
					}
				}
				if isStyleSheet && link != "" {
					links = append(links, link)
				}
			} else {
				for _, a := range n.Attr {
					if a.Key == attr {
						links = append(links, a.Val)
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
