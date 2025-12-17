package sort

// Algorithm overview:
// Loop through data starting from 2nd index
// For each element:
// - Iterate backwards, swapping each item unless current item is more than previous item

func InsertSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		j := i

		for j > 0 && data[j] < data[j-1] {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}

	return data
}