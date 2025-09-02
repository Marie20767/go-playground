package movezeroes

func MoveZeroes(nums []int) []int {
	reordered := make([]int, 0, len(nums))
	zeroCount := 0

	for _, num := range nums {
		if num != 0 {
			reordered = append(reordered, num)
		} else {
			zeroCount++
		}
	}

	return append(reordered, make([]int, zeroCount)...)
}

func MoveZeroesInPlace(nums []int) {
	i := 0

	for _, num := range nums {
		if num != 0 {
			nums[i] = num
			i++
		}
	}

	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
}
