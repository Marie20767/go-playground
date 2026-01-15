package palindrome

import (
	"regexp"
	"strings"
)

// Given a string s, return true if it is a palindrome, otherwise return false.

// A palindrome is a string that reads the same forward and backward. It is also case-insensitive and ignores all non-alphanumeric characters.

// Note: Alphanumeric characters consist of letters (A-Z, a-z) and numbers (0-9).

// Examples
//____________________________________
// Example 1:

// Input: s = "Was it a car or a cat I saw?"
// Output: true
// Explanation: After considering only alphanumerical characters we have "wasitacaroracatisaw", which is a palindrome.

// aba

// Example 2:
// Input: s = "tab a cat"
// Output: false

func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	var alphanumericRegex = regexp.MustCompile("[a-zA-Z0-9]")

	for i < j {
		if !alphanumericRegex.MatchString(string(s[i])) {
			i++
			continue
		}

		if !alphanumericRegex.MatchString(string(s[j])) {
			j--
			continue
		}

		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}
