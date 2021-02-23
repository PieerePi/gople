package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateLeft(s []int, num int) {
	if num <= 0 || num >= len(s) {
		return
	}
	reverse(s[:num])
	reverse(s[num:])
	reverse(s)
}

func rotateLeft2(s []int, num int) {
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
	rotateLeft(s, 0)
	fmt.Println(s) // [1 2 3 4 5]
	rotateLeft(s, 1)
	fmt.Println(s) // [2 3 4 5 1]
	rotateLeft(s, 2)
	fmt.Println(s) // [4 5 1 2 3]
	rotateLeft(s, 3)
	fmt.Println(s) // [2 3 4 5 1]
	rotateLeft(s, 4)
	fmt.Println(s) // [1 2 3 4 5]
	rotateLeft(s, 5)
	fmt.Println(s) // [1 2 3 4 5]
	rotateLeft(s, 6)
	fmt.Println(s) // [1 2 3 4 5]
}
