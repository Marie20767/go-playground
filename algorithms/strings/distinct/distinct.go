package distinct

// Find the largest 'distinct' substring from a given string.
// Where 'distinct' means containing unique characters only.

// E.g:
// Input: 'aaxyzbb'
// Output: 'axyzb'

func getSubstrings(s string, c int) []string {
	var substrings []string

	for i := 0; i < len(s)-c; i++ {
		substrings = append(substrings, s[i:c+i])
	}

	return substrings

}

func isDistinct(s string) bool {
	letterMap := make(map[string]struct{})

	for _, char := range s {
		if _, exists := letterMap[string(char)]; exists {
			return false
		}

		letterMap[string(char)] = struct{}{}
	}

	return true
}

func LargestDistinct(s string) string {
	if isDistinct(s) {
		return s
	}

	for i := 0; i < len(s); i++ {
		substrings := getSubstrings(s, len(s)-i)

		for _, substring := range substrings {
			if isDistinct(substring) {
				return substring
			}
		}
	}

	return s
}
