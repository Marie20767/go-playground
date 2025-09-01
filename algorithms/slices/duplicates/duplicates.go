package duplicates

func GetDuplicateOccurrences(nums []int) int {
	// [1, 1, 2, 3, 3, 2, 1] --> return 2

	duplicateCounter := 0
	numToOcurrences := make(map[int]int)

	for _, num := range nums {
		numToOcurrences[num]++
	}

	for _, occurrence := range numToOcurrences {
		if occurrence == 2 {
			duplicateCounter++
		}
	}

	return duplicateCounter
}