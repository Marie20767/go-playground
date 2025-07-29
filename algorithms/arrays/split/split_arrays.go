package split

func SplitTo3Elements(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var split [][]int
	currentChunk := []int{}

	for i, num := range nums {
		currentChunk = append(currentChunk, num)

		// Either the chunk is full or this is the last element
		if len(currentChunk) == 3 || i == len(nums)-1 {
			split = append(split, currentChunk)
			currentChunk = []int{}
		}
	}

	return split
}

// func SplitTo3Chunks(nums []int) [][]int {

// }

// new algo:
// - split array into multiple arrays in chunks of 3, return those chunks as elements in a new array

// example input:
// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]

// example output:
// [[1, 2, 3], [4, 5, 6], [7, 8, 9], [10, 11, 12]]

// if the last one doesn't fit into 3, return however many are remaining in the last chunk e.g:
// [1, 2, 3, 4] should output [[1, 2, 3], [4]]
