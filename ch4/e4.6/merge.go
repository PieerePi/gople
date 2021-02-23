package merge

import (
	"unicode"
	"unicode/utf8"
)

// MergeSpace removes adjacent repeated Unicode spaces from a byte slice.
func MergeSpace(s []byte) []byte {
	isSpace := false
	i, j := 0, 0
	for i < len(s) {
		r, size := utf8.DecodeRune(s[i:])
		if r == utf8.RuneError {
			panic("Rune Error")
		}
		if unicode.IsSpace(r) {
			if !isSpace {
				s[j] = ' '
				j++
			}
			isSpace = true
		} else {
			isSpace = false
			copy(s[j:j+size], s[i:i+size])
			j += size
		}
		i += size
	}
	return s[:j]
}
