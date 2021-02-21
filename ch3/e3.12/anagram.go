package main

func isAnagram(a, b string) bool {
	aFreq := make(map[rune]int)
	bFreq := make(map[rune]int)

	for _, c := range a {
		aFreq[c]++
	}
	for _, c := range b {
		bFreq[c]++
	}

	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}

	return true
}
