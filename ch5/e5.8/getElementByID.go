// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: getElementByID url id")
		os.Exit(1)
	}

	if err := getElementByID(os.Args[1], os.Args[2]); err != nil {
		fmt.Printf("getElementByID failed: %v\n", err)
		os.Exit(1)
	}
}

func getElementByID(url, id string) error {
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

	node := ElementByID(doc, id)
	if node == nil {
		fmt.Printf("No element with id %q found\n", id)
	} else {
		forEachNode(node, printStartElement, printEndElement)
	}

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
// pre/post/forEachNode return value:
//    true to stop, false to continue
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if pre(n) {
			return true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if forEachNode(c, pre, post) {
			return true
		}
	}

	if post != nil {
		if post(n) {
			return true
		}
	}

	return false
}

//!-forEachNode

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node

	forEachNode(doc, func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				node = n
				return true
			}
		}

		return false
	}, nil)

	return node
}

var depth int

func printStartElement(n *html.Node) bool {
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
	return false
}

func printEndElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	return false
}
