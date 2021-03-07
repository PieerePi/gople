package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("change foo to upper case, that is $foo", toUpper))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}
