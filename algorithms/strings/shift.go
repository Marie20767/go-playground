package shift

func Shift(s string) string {
	shifted := []rune(s)

	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			if char == 'z' {
				shifted[i] = 'a'
			} else {
				shifted[i] = char + 1
			}
		}
	}

	return string(shifted)
}
