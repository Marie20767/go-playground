package arrays

// given an array of numbers:
// - start at the first one, it represents how many steps you could jump (e.g. if it's 3 you can jump forward 1, 2 or 3 steps)
// should return true if it's possible to land at the last index of the array with some sequence of jumps

func CanJump(nums []int) bool {
	maxReach := 0
	for i, jump := range nums {
		if i > maxReach {
			return false
		}

		if i+jump > maxReach {
			maxReach = i + jump
		}
		
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	
	return true
}
