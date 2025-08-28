package shift

import (
	"strings"
)


func Shift(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	shifted := []rune(s)

	for i, char := range s {
		currentCharI := strings.Index(alphabet, string(char))

		if currentCharI >= len(alphabet)-1 {
			shifted[i] = rune(alphabet[0])
		} else {
			shifted[i] = rune(alphabet[currentCharI+1])
		}
	}

	return string(shifted)
}
