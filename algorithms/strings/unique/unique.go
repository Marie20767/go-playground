package unique

func FirstUnique(s string) int {
	occurrences := make(map[rune]int)

	for _, char := range s {
		occurrences[char]++
	}

	runeIndex := 0

	// when ranging over string, i is byte index not a position index!
	for _, char := range s {
		if occurrences[char] == 1 {
			return runeIndex
		}

		runeIndex++
	}

	return -1
}