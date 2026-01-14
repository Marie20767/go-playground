package subsets

// Given an array nums of unique integers, return all possible subsets of nums.
// The solution set must not contain duplicate subsets. You may return the solution in any order.

// Examples
// _________________________________________________

// Example 1:
// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Example 2:
// Input: nums = [7]
// Output: [[],[7]]

func Subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(i int)

	subset := []int{}
	dfs = func(i int) {
		if i >= len(nums) {
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			res = append(res, subsetCopy)
			return
		}

		// decision to include nums[i]
		subset = append(subset, nums[i])
		dfs(i + 1)

		// decision NOT to include nums[i]
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)
	return res
}