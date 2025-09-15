package slidingmaximum

// Example Input: [1,3,-1,-3,5,3,6,7]
// Example Output: [3,3,-1,5,5,6,7]

func SlidingMax(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	var output []int

	for i := 0; i < len(n)-1; i++ {
		if n[i] > n[i+1] {
			output = append(output, n[i])
		} else {
			output = append(output, n[i+1])
		}
	}

	return output
}
