package split_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays/split"
	"github.com/stretchr/testify/assert"
)

func TestSplitArrays(t *testing.T) {
	t.Run("should split array in 3 chunks", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

		expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
		actual := split.SplitTo3Elements(nums)

		assert.Equal(t, expected, actual)
	})

	t.Run("should split array in 2 chunks", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}

		expected := [][]int{{1, 2, 3}, {4}}
		actual := split.SplitTo3Elements(nums)

		assert.Equal(t, expected, actual)
	})
}
