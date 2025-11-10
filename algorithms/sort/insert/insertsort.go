package insertsort

// Algorithm overview:
// Loop through data starting from 2nd index (1st element is already a sorted sub-array)
// For each element:
// - Iterate backwards, swapping each item unless current item is more than previous item
func InsertionSort(data []int) []int {
	// 7, 5, 3, 6, 1, 0 // i = 1, j = 1 {5, 1, 7}

	for i := 1; i < len(data); i++ {
		j := i

		for j > 0 && data[j] < data[j-1] {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}

	return data
}
