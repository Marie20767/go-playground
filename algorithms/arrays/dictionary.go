package arrays

import (
	"sort"
	"strings"
)

// CharRemoval(strArr) reads the array of strings stored in strArr, which will contain 2 elements:
//   - 1st element: sequence of chars representing a word
//   - 2nd element: dictionary (i.e. string of comma-separated words in alphabetical order)
//     For example: ["worlcde", "apple,bat,cat,goodbye,hello,yellow,why,world"]
// Determine the minimum number of characters, if any, that can be removed from the word so that it matches one of the words from the dictionary
// If the word cannot be found no matter what characters are removed, return -1.
// CharRemoval(["worlcde", "apple,bat,cat,goodbye,hello,yellow,why,world"]) returns 2 because once you remove the "c" and "e" you are left with "world" which exists within the dictionary

func CharRemoval(strArr []string) int {
	dict := strings.Split(strArr[1], ",")
	target := strArr[0]
	sort.Slice(dict, func(i, j int) bool {
		return len(dict[i]) > len(dict[j])
	})

	for _, word := range dict {
		if isSubsequence(target, word) {
			return len(target) - len(word)
		}
	}

	return -1
}

func isSubsequence(target, word string) bool {
		i := 0

		for j := 0; j < len(target) && i < len(word); j++ {
			if target[j] == word[i] {
				i++
			}
		}

		return i == len(word)
}