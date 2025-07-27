package merge

func SortMergedNums(nums1 []int, nums2 []int) []int {
	nums2ToCheck := make([]int, len(nums2))
	copy(nums2ToCheck, nums2)

	merged := make([]int, 0, len(nums1) + len(nums2))

	// TODO: counter


	// nums1 := []int{5, 10, 50, 9000}
	// nums2 := []int{1, 7, 21}

	for i := 0; i < len(nums1); i++ {
		if len(nums2ToCheck) == 0 {
			merged = append(merged, nums1[i])
		} else {
			for j := 0; j < len(nums2ToCheck); j++ {
				if nums1[i] > nums2ToCheck[j] {
					merged = append(merged, nums2ToCheck[j])
					nums2ToCheck = append(nums2ToCheck[:i], nums2ToCheck[i+1:]...)
				} else {
					merged = append(merged, nums1[i])
					break;
				}
			}
		}
	}

	return merged
}




// nums1 := []int{1, 20, 100, 500}   j 1
// nums2 := []int{2, 3, 4, 5, 6}   i 4
// result:= []

// rule for stopping => i == len(nums1) || j == len(nums2)
// while loop

// func(nums1, nums2)
// 		i := 0 // counter nums1
// 	 j := 0 // counter nums2
//	merged := []
// 	while i == len(nums1) || j == len(nums2):
//		if nums1[i] < nums2[j] {
//			merged[i + j] =  num1[i]
//			i++
//		}
//    else {
//		 	merged[i + j] = nums2[j]
//			j++
//		}
// check which nums array hasn't reached the end e.g. nums1
// add  all the nums of nums1 into merged
// use for loop not append