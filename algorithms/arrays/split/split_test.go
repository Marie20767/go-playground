package split_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays/split"
	"github.com/stretchr/testify/assert"
)

func TestSplitTo3Elements(t *testing.T) {
	t.Run("should split slice into chunks of 3 elements each", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

		expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
		actual := split.SplitTo3Elements(nums)

		assert.Equal(t, expected, actual)
	})

	t.Run("should split slice into chunks of up to 3 elements", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}

		expected := [][]int{{1, 2, 3}, {4}}
		actual := split.SplitTo3Elements(nums)

		assert.Equal(t, expected, actual)
	})
}

func TestSplitTo3Chunks(t *testing.T) {
	t.Run("should split slice into 3 equal chunks", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

		expected := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
		actual := split.SplitTo3Chunks(nums)

		assert.Equal(t, expected, actual)
	})

	t.Run("should split slice into 3 unequal chunks", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		expected := [][]int{{1, 2}, {3, 4}, {5}}
		actual := split.SplitTo3Chunks(nums)

		assert.Equal(t, expected, actual)
	})
}
