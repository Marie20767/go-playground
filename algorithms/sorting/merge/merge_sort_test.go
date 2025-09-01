package merge_sort_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/sorting/merge"
	"github.com/stretchr/testify/assert"
)

func TestSortMerged(t *testing.T) {
	t.Run("should merge and sort 2 slices", func(t *testing.T) {
		nums1 := []int{1, 20, 100, 500}
		nums2 := []int{2, 3, 4, 5, 6}

		expected := []int{1, 2, 3, 4, 5, 6, 20, 100, 500}
		actual := merge_sort.SortMerged(nums1, nums2)

		assert.Equal(t, expected, actual)
	})
}

func TestMergeSort(t *testing.T) {
	t.Run("should sort slice with merge sort", func(t *testing.T) {
		nums := []int{100, 5, 1, 1000, 10}
		expected := []int{1, 5, 10, 100, 1000}
		actual := merge_sort.MergeSort(nums)

		assert.Equal(t, expected, actual)
	})
}
