package merge_sort_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/sorting/merge"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	t.Run("should sort slice with merge sort", func(t *testing.T) {
		nums := []int{100, 5, 1, 1000, 10}
		expected := []int{1, 5, 10, 100, 1000}
		actual := merge_sort.MergeSort(nums)

		assert.Equal(t, expected, actual)
	})
}
