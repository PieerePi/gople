package main

import (
	"testing"
)

type PalindromeChecker []byte

func (x PalindromeChecker) Len() int           { return len(x) }
func (x PalindromeChecker) Less(i, j int) bool { return x[i] < x[j] }
func (x PalindromeChecker) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func TestIsPalindrome2(t *testing.T) {
	// fmt.Println(IsPalindrome(PalindromeChecker([]byte("abcdcba"))))
	// fmt.Println(IsPalindrome(PalindromeChecker([]byte("abcdecba"))))
	if !IsPalindrome(PalindromeChecker([]byte("abcdcba"))) {
		t.Errorf("IsPalindrome error")
	}
	if IsPalindrome(PalindromeChecker([]byte("abcdecba"))) {
		t.Errorf("IsPalindrome error")
	}
}
