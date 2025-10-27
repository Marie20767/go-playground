package runlength

import (
	"fmt"
)

// Have the function RunLength(str) take the str parameter being passed and return a compressed version of the string using the Run-length encoding algorithm.
// This algorithm works by taking the occurrence of each repeating character and outputting that number along with a single character of the repeating sequence.
// For example: "wwwggopp" would return 3w2g1o2p. The string will not contain any numbers, punctuation, or symbols.

func Runlength(s string) string {
	compressed := ""

	if len(s) <= 0 {
		return compressed
	}

	occurrence := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			occurrence++
		} else {
			compressed += fmt.Sprintf("%d%s", occurrence, string(s[i-1]))
			occurrence = 1
		}
	}

	compressed += fmt.Sprintf("%d%s", occurrence, string(s[len(s)-1]))

	return compressed
}
