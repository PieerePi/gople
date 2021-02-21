package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"aba", "baa", true},
		{"aaa", "baa", false}, // same characters but different frequencies
		{"你好世界", "世界你好", true},
		{"你好世界", "世界您好", false},
	}
	for _, test := range tests {
		got := isAnagram(test.a, test.b)
		if got != test.want {
			t.Errorf("isAnagram(%q, %q), got %v, want %v",
				test.a, test.b, got, test.want)
		}
	}
}
