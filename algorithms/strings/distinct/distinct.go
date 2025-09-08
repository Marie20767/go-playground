package distinct

// Find the largest 'distinct' substring from a given string.
// Where 'distinct' means containing unique characters only.

// E.g:
// Input: 'aaxyzbb'
// Output: 'axyzb'

// distinct(s string) string {} go through every possible substring and find largest distinct substring
//      distinct = ""
// 		- substrings(s string) []string {} takes a string and returns all possible substrings
//         - loop through s
//         - substrings []string = [] + s[i]
//         - loop through every letter after s[i]
//         - substrings = s[i] + s[j]
//         - return substrings
// 		- isDistinct(s string) bool {} takes a substring and return s == distinct
//         - make map[string]:{}
//         - loop through s
//         - if map[s[i]] return false else append to map
//   - longestUniqueSubstring:
//          - loop through every substring
//          - if distinct(s[i]) == true && len(s[i]) > len(distinct) { distinct = substring }
//          - return distinct

func getSubstrings(s string) []string {
	var substrings []string

	for i := 0; i < len(s); i++ {
		substrings = append(substrings, string(s[i]))
		prevSubstring := string(s[i])

		for j := i + 1; j < len(s); j++ {
			newSubstring := prevSubstring + string(s[j])
			substrings = append(substrings, newSubstring)
			prevSubstring = newSubstring
		}
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
	distinct := ""

	for _, s := range getSubstrings(s) {
		if isDistinct(s) && len(s) > len(distinct) {
			distinct = s
		}
	}

	return distinct
}
