package merge

import (
	"unicode"
	"unicode/utf8"
)

func MergeSpace(s []byte) []byte {
	isSpace := false
	i, j := 0, 0
	for i < len(s) {
		r, size := utf8.DecodeRune(s[i:])
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
