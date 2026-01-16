package twointegersum

// Given an array of integers numbers that is sorted in non-decreasing order.
// Return the indices (1-indexed) of two numbers, [index1, index2], such that they add up to a given target number target and index1 < index2.
// Note that index1 and index2 cannot be equal, therefore you may not use the same element twice.
// There will always be exactly one valid solution.

// Examples
// ______________________
// Example 1:
// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
// Example 2:

// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

// Example 3:
// Input: numbers = [-1,0], target = -1
// Output: [1,2]
// Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

func TwoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		curSum := numbers[i] + numbers[j]
		switch {
		case curSum > target:
			j--
		case curSum < target:
			i++
		default:
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}
