// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
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

	//!+call
	var depth int
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.CommentNode {
			fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
		} else if n.Type == html.TextNode {
			trimmed := strings.TrimSpace(n.Data)
			if trimmed != "" {
				// for correct indentation of multiple lines, print each line separately
				lines := strings.Split(trimmed, "\n")
				for _, l := range lines {
					fmt.Printf("%*s%s\n", depth*2, "", l)
				}
			}
		} else if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, a := range n.Attr {
				fmt.Printf(" %s=%q", a.Key, a.Val)
			}
			if n.FirstChild != nil {
				fmt.Printf(">\n")
			} else {
				fmt.Printf(" />\n")
			}
			depth++
		}
	}, func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			if n.FirstChild != nil {
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		}
	})
	//!-call

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
