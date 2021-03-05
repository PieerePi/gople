// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	order := topoSort(prereqs)
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	valid := isValidTopoSort(order)
	if valid {
		fmt.Printf("This order is a valid toposort.\n")
	} else {
		fmt.Printf("This order is not a valid toposort.\n")
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for key := range m {
		visitAll([]string{key})
	}
	return order
}

func isValidTopoSort(order []string) bool {
	orderMap := make(map[string]int)

	for i, course := range order {
		orderMap[course] = i
	}

	for course, i := range orderMap {
		for _, prereq := range prereqs[course] {
			if i < orderMap[prereq] {
				return false
			}
		}
	}
	return true
}

//!-main
