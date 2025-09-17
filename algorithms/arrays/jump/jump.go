package jump

// given an array of numbers:
// - start at the first one, it represents how many steps you could jump (e.g. if it's 3 you can jump forward 1, 2 or 3 steps)
// should return true if it's possible to land at the last index of the array with some sequence of jumps

	func CanJump(n []int) bool {
		if len(n) == 0 {
			return false
		}
		
		if n[0] <= 0 {
			return false
		}

		i := 0

		for range n {
			if i+n[i] >= len(n)-1 {
				return true
			}

			if n[n[i]+i] == 0 {
				n[i] = n[i] - 1
				return CanJump(n)
			}

			i += n[i]
		}

		return false
	}
