package merge_sort_test

import (
	"testing"

	merge_sort "github.com/Marie20767/go-playground/algorithms/sorting/merge"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	t.Run("should sort merged numbers into ascending order", func(t *testing.T) {
		nums1 := []int{1, 20, 100, 500}
		nums2 := []int{2, 3, 4, 5, 6}

		expected := []int{1, 2, 3, 4, 5, 6, 20, 100, 500}
		actual := merge_sort.SortMergedNums(nums1, nums2)

		assert.Equal(t, expected, actual)
	})
}
