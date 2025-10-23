package slidingwindowmax

// So if you have something like:
// Input: [1,3,-1,-3,5,3,6,7]  AND  k=4
// Then you'd first be comparing 1, 3, -1, -3
// Then you'd compare 3, -1, -3, 5
// Then -1, -3, 5, 3
// etc.
// Example Output would be: [3,5,5,6,7]

func getHighestNum(section []int) int {
	highestNum := section[0]

	for i := 1; i < len(section); i++ {
		if section[i] > highestNum {
			highestNum = section[i]
		}
	}

	return highestNum
}

func SlidingWindowMax(n []int, k int) []int {
	if len(n) <= 1 {
		return n
	}

	var output []int

	if k > len(n) {
		output = append(output, getHighestNum(n))
		return output
	}

	i := 0

	for k+i <= len(n) {
		section := n[i : k+i]

		output = append(output, getHighestNum(section))
		i++
	}

	return output
}
