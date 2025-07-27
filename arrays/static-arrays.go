package main

import (
	"fmt"
	"time"
)

func main() {
	// nums := []int{1, 1, 2, 3, 4}
	// k := removeDuplicatesWithExtraMemory(nums)
	// k := removeDuplicates(nums)
	// k := removeElementsWithExtraMemory(nums, 1)
	// k := removeElements(nums, 1)

	// fmt.Println("k", k)
	// fmt.Println("nums", nums[:k])

	// a := []int{1, 2, 3, 4, 5, 6}
	// const i, j = 2, 4
	// // a = append(a[:i], a[j+1:]...) // [1, 2, 6]
	// a = append(a[:i], a[j:]...) // [1, 2, 5, 6]

	// fmt.Println("a", a)

	// nums := []int{1, 1, 2, 3, 3, 2}
	// fmt.Println("occurrences", getDuplicateOccurrences(nums))

	// fmt.Println("hols exit balance", holidayExitBalance(25, 10, "28/09/2025"))

	// nums := []int{8, 2, 3, 1, 5, 0}

	// fmt.Println("sorted nums", sortedNumsBubbleSort(nums))
	nums1 := []int{5, 10, 50, 9000}
	nums2 := []int{1, 7, 21}
	fmt.Println("merged nums", sortMergedNums(nums1, nums2))
}

func removeDuplicatesWithExtraMemory(nums []int) int {
	seen := make(map[int]bool)          // map[]
	unique := make([]int, 0, len(nums)) //[]

	// _, means ignore the index
	// range is a special keyword for iterating over collections: slices, arrays, maps, strings, and channels
	// range unpacks the index/key and value during the loop.
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			unique = append(unique, num)
		}
	}

	// if you don't use range:
	// for i := 0; i < len(nums); i++ {
	//   ...
	// }

	copy(nums, unique)

	return len(unique)
}

func removeDuplicates(nums []int) int {
	k := 0

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[k] {
			k++
			nums[k] = nums[j]
		}
	}

	return k + 1
}

func removeElementsWithExtraMemory(nums []int, val int) int {
	temp := make([]int, 0, len(nums))

	for _, num := range nums {
		if num != val {
			temp = append(temp, num)
		}
	}

	copy(nums, temp)

	return len(temp)
}

func removeElements(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}

func pushLastUuidToFront(ids []string) []string {
	// assume that all ids are unique
	if len(ids) < 0 {
		return ids
	}

	lastId := ids[len(ids)-1]
	idsWithoutLastId := append(ids[:len(ids)-1])
	reorderedIds := append([]string{lastId}, idsWithoutLastId...)

	return reorderedIds
}

func pushLastUuidToFrontInPlace(ids []string) []string {
	// assume that all ids are unique
	if len(ids) < 0 {
		return ids
	}

	originalIds := make([]string, len(ids))
	copy(originalIds, ids)

	for i := 0; i < len(ids); i++ {
		if i == len(ids)-1 {
			ids[0] = originalIds[i]
		} else {
			ids[i+1] = originalIds[i]
		}
	}

	return ids
}

func getDuplicateOccurrences(nums []int) int {
	// [1, 1, 2, 3, 3, 2, 1] --> return 2

	duplicateCounter := 0
	numToOcurrences := make(map[int]int)

	for _, num := range nums {
		numToOcurrences[num]++
	}

	for _, occurrence := range numToOcurrences {
		if occurrence == 2 {
			duplicateCounter++
		}
	}

	return duplicateCounter
}

// Fresto works at Grooble, given T is the total number of holidays that she gets per year,
// R is the remaning number of holidays she has for the current year and D is the date she is planning to leave
// (as a string in DD/MM/YYYY format), write a function that takes T, R and D and returns the number of holidays that Fresto is due in refund or has in excess.
// E.g. if she leaves exactly halfway through the year and has used exactly half her holiday, the function should return 0.

func holidayExitBalance(t int, r int, d string) float32 {
	const workingDaysPerYear = 260
	const workingDaysPerMonth = 21.6
	const dateTimeLayout = "02/01/2006"

	datetime, err := time.Parse(dateTimeLayout, d)

	if err != nil {
		panic(err)
	}

	// t = 25
	// r = 10
	// d = 28/09/2025

	dailyHolAllowance := float32(t) / workingDaysPerYear

	_, month, day := datetime.Date()
	monthNum := int(month)

	totalDaysWorkedUpToPrevMonth := float32(monthNum-1) * workingDaysPerMonth
	totalDaysWorkedLeavingMonth := float32(day) / 30 * workingDaysPerMonth
	totalDaysWorked := totalDaysWorkedUpToPrevMonth + totalDaysWorkedLeavingMonth

	holAllowanceUntilExit := dailyHolAllowance * float32(totalDaysWorked)

	return holAllowanceUntilExit - float32(t-r)
}

func sortedNumsBubbleSort(nums []int) []int {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return nums
}

func sortMergedNums(nums1 []int, nums2 []int) []int {
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