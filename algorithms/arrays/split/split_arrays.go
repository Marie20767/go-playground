package split

import "math"

func SplitTo3Elements(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var split [][]int
	currentChunk := []int{}

	for i, num := range nums {
		currentChunk = append(currentChunk, num)

		// either the chunk is full or this is the last element
		if len(currentChunk) == 3 || i == len(nums)-1 {
			split = append(split, currentChunk)
			currentChunk = []int{}
		}
	}

	return split
}

func SplitTo3Chunks(nums []int) [][]int {

	numElementsPerSlice := int(math.Ceil(float64(len(nums)) / 3))

	var split [][]int
	chunk := []int{}

	for i, num := range nums {
		chunk = append(chunk, num)

		if len(chunk) == numElementsPerSlice || i == len(nums)-1 {
			split = append(split, chunk)
			chunk = []int{}
		}
	}

	return split
}
