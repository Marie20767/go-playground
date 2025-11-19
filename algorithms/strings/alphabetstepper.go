package strings

func AlphabetStepper(s string) string {
	shifted := []rune(s)

	for i, char := range s {
		switch {
		case char >= 'a' && char <= 'z':
			if char == 'z' {
				shifted[i] = 'a'
			} else {
				shifted[i] = char + 1
			}
		case char >= 'A' && char <= 'Z':
			if char == 'Z' {
				shifted[i] = 'A'
			} else {
				shifted[i] = char + 1
			}
		}
	}

	return string(shifted)
}