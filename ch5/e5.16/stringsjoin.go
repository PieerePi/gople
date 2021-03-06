package main

import (
	"bytes"
	"fmt"
	"strings"
)

func stringsJoin(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}

func stringsJoin2(sep string, elems ...string) string {
	var buf bytes.Buffer
	if len(elems) == 0 {
		return ""
	}
	buf.WriteString(elems[0])
	for _, elem := range elems[1:] {
		buf.WriteString(sep)
		buf.WriteString(elem)
	}
	return buf.String()
}

func main() {
	fmt.Println(strings.Join([]string{}, " xx "))
	fmt.Println(strings.Join([]string{"ab"}, " xx "))
	fmt.Println(strings.Join([]string{"ab", "cd"}, " xx "))
	fmt.Println(stringsJoin(" xx "))
	fmt.Println(stringsJoin(" xx ", "ab"))
	fmt.Println(stringsJoin(" xx ", "ab", "cd"))
	fmt.Println(stringsJoin2(" xx "))
	fmt.Println(stringsJoin2(" xx ", "ab"))
	fmt.Println(stringsJoin2(" xx ", "ab", "cd"))
}
