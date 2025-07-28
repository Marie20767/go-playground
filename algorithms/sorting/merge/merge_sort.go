package merge_sort

func SortMergedNums(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return append([]int{}, nums2...)
	}
	
	if len(nums2) == 0 {
		return append([]int{}, nums1...)
	}
	
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		merged = append(merged, nums1[i])
		i++
	}

	for j < len(nums2) {
		merged = append(merged, nums2[j])
		j++
	}

	return merged
}
