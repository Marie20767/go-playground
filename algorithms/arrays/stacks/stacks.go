package stacks

import (
	"github.com/Marie20767/go-playground/data_structures/stack"
)

func IsStackValid(s string) bool {
	closeToOpenBracketsMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := stack.NewStack()

	for _, c := range s {
		if openBracket, exists := closeToOpenBracketsMap[c]; exists {
			if !stack.Empty() {
				top, ok := stack.Pop()

				if ok && top != openBracket {
					return false
				}
			} else {
				return false
			}
		} else {
			stack.Push(c)
		}
	}

	return stack.Empty()
}
