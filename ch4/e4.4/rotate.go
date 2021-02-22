package main

import (
	"fmt"
)

func roateLeft(s []int, num int) {
	slen := len(s)
	if num <= 0 || num >= slen {
		return
	}
	left := make([]int, num)
	copy(left, s[:num])
	copy(s, s[num:])
	copy(s[slen-num:], left)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	roateLeft(s, 0)
	fmt.Println(s) // [1 2 3 4 5]
	roateLeft(s, 1)
	fmt.Println(s) // [2 3 4 5 1]
	roateLeft(s, 2)
	fmt.Println(s) // [4 5 1 2 3]
	roateLeft(s, 3)
	fmt.Println(s) // [2 3 4 5 1]
	roateLeft(s, 4)
	fmt.Println(s) // [1 2 3 4 5]
	roateLeft(s, 5)
	fmt.Println(s) // [1 2 3 4 5]
	roateLeft(s, 6)
	fmt.Println(s) // [1 2 3 4 5]
}
