package arrays

import "slices"

// Group Anagrams

// Input: list of words
// Output: lists where each group contains anagrams contained within the input (i.e. an array of arrays)

// Example:
// input: ['deed', 'edde', 'bye', 'hello', 'ybe']
// output:
// [ ['deed', 'edde'], ['bye', 'ybe'] ]

// If no anagrams at all => can return empty array

func sortWord(word string) string {
	letters := []rune(word)
	slices.Sort(letters)
	return string(letters)
}

func Anagrams(words []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		key := sortWord(word)
		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0)
	
	for _, group := range groups {
		if len(group) > 1 {
			result = append(result, group)
		}
	}

	return result
}
