package peakfinder

func PeakFinder(n []int) []int {
	peaks := make([]int, 0, len(n))

	if len(n) <= 1 {
		return peaks
	}

	for i, num := range n {
		left := -1 << 63
		right := -1 << 63

		if i > 0 {
			left = n[i - 1]
		}

		if i < len(n) - 1 {
			right = n[i + 1]
		}

		if num > left && num > right {
			peaks = append(peaks, num)
		}
	}

	return peaks
}
