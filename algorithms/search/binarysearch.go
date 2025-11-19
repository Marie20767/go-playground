package search

// Algorithm overview:
// Takes an array (can be of ints for our implementation) and a target value that we need to find in the array
// Returns the index in the array of the target value OR -1 if not found alongside a bool (false if -1)

// Input must be in sorted order already.
// Find the value at midpoint, then:
// - if target equals the value => return index of this value
// - if the target is less => find value at midpoint on left side => ...
// - if the target is more => find value at midpoint on right side => ...

func BinarySearch(n []int, target int) (int, bool) {
	start, end := 0, len(n)-1

	for start <= end {
		midpoint := (end-start)/2 + start

		switch {
		case n[midpoint] == target:
			return midpoint, true
		case target > n[midpoint]:
			start = midpoint + 1
		case target < n[midpoint]:
			end = midpoint - 1
		}
	}

	return -1, false
}

func BinarySearchRecursive(n []int, target int) (int, bool) {
	return binarySearchHelper(n, target, 0, len(n)-1)
}

func binarySearchHelper(n []int, target, start, end int) (int, bool) {
	midpoint := (end-start)/2 + start

	if start > end {
		return -1, false
	}

	if n[midpoint] == target {
		return midpoint, true
	}

	if target > n[midpoint] {
		start = midpoint + 1
		return binarySearchHelper(n, target, start, end)
	}

	end = midpoint - 1
	return binarySearchHelper(n, target, start, end)
}