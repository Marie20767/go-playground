package split

func SplitTo3Elements(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var split [][]int

	currentChunk := 0
	counter := 0

	for _, num := range nums {
		if counter == 3 {
			currentChunk++
			counter = 0
		}

		if counter == 0 {
			split = append(split, []int{})
		}

		if counter < 3 {
			split[currentChunk] = append(split[currentChunk], num)
			counter++
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
