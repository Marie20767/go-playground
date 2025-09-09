package ransom

// Input: note (a string), magazine (a string)

// Output: True/False if note can be built from magazine letters (case-sensitive, each letter used at most once).

// Example:
// ransomNote("hello", "ohelp") => false
// ransomNote("hello", "yellowh") => true
// ransomNote("hello", "Hello") => false


func containsAllLetters(noteMap map[string]int, magazineMap map[string]int) bool {
	for l, c := range noteMap {
		if magazineMap[l] < c {
			return false
		}
	}
	
	return true
}

func Ransom(note string, magazine string) bool {
	if len(magazine) < len(note) {
		return false
	}

	noteMap := make(map[string]int)

	for _, d := range note {
		noteMap[string(d)]++
	}

	magazineMap := make(map[string]int)

	for _, d := range magazine {
		if _, exists := noteMap[string(d)]; exists  {
			magazineMap[string(d)]++
		}
	}

	return containsAllLetters(noteMap, magazineMap)
}
