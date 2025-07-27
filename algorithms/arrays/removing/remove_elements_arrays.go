package removing

func RemoveDuplicatesWithExtraMemory(nums []int) int {
	seen := make(map[int]bool)          // map[]
	unique := make([]int, 0, len(nums)) //[]

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			unique = append(unique, num)
		}
	}

	copy(nums, unique)

	return len(unique)
}

func RemoveDuplicates(nums []int) int {
	k := 0

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[k] {
			k++
			nums[k] = nums[j]
		}
	}

	return k + 1
}

func RemoveElementsWithExtraMemory(nums []int, val int) int {
	temp := make([]int, 0, len(nums))

	for _, num := range nums {
		if num != val {
			temp = append(temp, num)
		}
	}

	copy(nums, temp)

	return len(temp)
}

func RemoveElements(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}