package strings

import "strings"

func Reverse(s string) string {
	words := strings.Fields(strings.TrimSpace(s))
	reversed := make([]string, 0, len(words))

	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}


	return strings.Join(reversed, " ")
}