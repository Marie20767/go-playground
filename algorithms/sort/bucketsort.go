package sort

import (
	"math"
)

// Version 1:
// Put elements into a series of buckets that hold a specific value (e.g. 5 buckets: 10, 20, 30, 40, 50)
// - Concatenate those buckets (or replace values of the original array in place)
// - Note - numberOfBuckets could be constant or could be calculated as part of the algorithm

func BucketSort(n []int) {
	fixedBucketsMap := map[int]int{
		10: 0,
		20: 1,
		30: 2,
		40: 3,
		50: 4,
	}

	sorted := make([][]int, len(fixedBucketsMap))

	for _, num := range n {
		index, _ := fixedBucketsMap[num]
		sorted[index] = append(sorted[index], num)
	}

	currentIndex := 0

	for _, bucket := range sorted {
		for _, num := range bucket {
			n[currentIndex] = num
			currentIndex++
		}
	}
}

// Version 2:
// Put elements into a series of buckets that hold a specific range (e.g. 5 buckets 0-10, 11-20, etc)
// - Sort those buckets separately (e.g. with insertion sort)
// - Concatenate those buckets (or replace values of the original array in place)

func BucketSortInRange(n []int) {
	if len(n) <= 0 {
		return
	}

	numBuckets := int(math.Sqrt(float64(len(n))))
	min, max := findMinMax(n)
	bucketRange := float64((max - min + 1)) / float64(numBuckets)
	buckets := make([][]int, numBuckets)

	for _, num := range n {
		index := int(float64((num - min)) / bucketRange)

		if index >= numBuckets {
			index = numBuckets - 1
		}

		buckets[index] = append(buckets[index], num)
	}

	currentIndex := 0

	for _, bucket := range buckets {
		sorted := InsertSort(bucket)

		for _, num := range sorted {
			n[currentIndex] = num
			currentIndex++
		}
	}
}

func findMinMax(n []int) (int, int) {
	min := n[0]
	max := n[0]

	for i := 1; i < len(n); i++ {
		switch {
		case n[i] > max:
			max = n[i]
		case n[i] < min:
			min = n[i]
		}
	}

	return min, max
}
