package unique

// UniqueString removes adjacent repeated strings from a string slice.
func UniqueString(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}
	last := 0
	for _, s := range strs {
		if strs[last] == s {
			continue
		}
		last++
		strs[last] = s
	}
	return strs[:last+1]
}
