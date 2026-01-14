package combinationsum

// You are given an array of distinct integers nums and a target integer target. Your task is to return a list of all unique combinations of nums where the chosen numbers sum to target.
// The same number may be chosen from nums an unlimited number of times. Two combinations are the same if the frequency of each of the chosen numbers is the same, otherwise they are different.
// You may return the combinations in any order and the order of the numbers in each combination can be in any order.

// Examples
// __________________________________________
// Example 1:
// Input:
// nums = [2,5,6,9]
// target = 9
// Output: [[2,2,5],[9]]

// Example 2:
// Input:
// nums = [3,4,5]
// target = 16
// Output: [[3,3,3,3,4],[3,3,5,5],[4,4,4,4],[3,4,4,5]]

// Example 3:
// Input:
// nums = [3]
// target = 5
// Output: []

func CombinationSum(nums []int, target int) [][]int {
	res := [][]int{}

	var dfs func(int, []int, int)
	dfs = func(i int, subset []int, total int) {
		if total == target {
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			res = append(res, subsetCopy)
			return
		}

		if i >= len(nums) || total > target {
			return
		}

		subset = append(subset, nums[i])
		dfs(i, subset, total+nums[i])
		subset = subset[:len(subset)-1]
		dfs(i+1, subset, total)
	}

	dfs(0, []int{}, 0)
	return res
}
