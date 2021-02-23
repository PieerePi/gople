package anagram

// IsAnagram checks if two strings are anagram.
func IsAnagram(a, b string) bool {
	aFreq := make(map[rune]int)
	bFreq := make(map[rune]int)

	for _, c := range a {
		aFreq[c]++
	}
	for _, c := range b {
		bFreq[c]++
	}

	if len(aFreq) != len(bFreq) {
		return false
	}
	for k, av := range aFreq {
		if bv, ok := bFreq[k]; !ok || bv != av {
			return false
		}
	}

	return true
}
