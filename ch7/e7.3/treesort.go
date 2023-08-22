// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package main

// !+
import (
	"fmt"
	"math/rand"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-

// interface fmt.Stringer
func (t *tree) String() string {
	data := make([]int, 50)
	data = appendValues(data[:0], t)
	var str strings.Builder
	str.WriteString("[")
	for i := range data {
		// interface io.Writer
		if i == 0 {
			fmt.Fprintf(&str, "%d", data[i])
		} else {
			fmt.Fprintf(&str, " %d", data[i])
		}
	}
	str.WriteString("]")
	return str.String()
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	fmt.Println(root)
}
