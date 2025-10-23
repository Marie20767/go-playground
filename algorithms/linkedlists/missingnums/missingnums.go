package missingnums

func MissingNums(nums []int) int {
	numsMap := make(map[int]struct{})

	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, exists := numsMap[i]; !exists {
			return i
		}
	}

	return 0
}

