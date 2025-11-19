package arrays

import (
	"strings"
)

func AreBracketsValid(s string) bool {
	closeToOpenBracketsMap := map[string]string{
		")" : "(", 
		"]" : "[", 
		"}" : "{",
	}

	brackets := strings.Split(s, "")
	isValid := true

	for isValid && len(brackets) > 0 {
		closingBracket := brackets[len(brackets) - 1]
		openBracket := closeToOpenBracketsMap[closingBracket]

		if brackets[0] == openBracket {
			brackets = brackets[1: len(brackets) - 1]
		} else {
			isValid = false
		}
	}

	return isValid
}
