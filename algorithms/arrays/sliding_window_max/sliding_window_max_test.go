package slidingwindowmax_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays/sliding_window_max"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowMax(t *testing.T) {
	t.Run("Should return empty array", func(t *testing.T) {
		n := []int{}
		k := 3
		expected := []int{}
		actual := slidingwindowmax.SlidingWindowMax(n, k)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return max with 1 element when k > len(n)", func(t *testing.T) {
		n := []int{-100, -1000}
		k := 3
		expected := []int{-100}
		actual := slidingwindowmax.SlidingWindowMax(n, k)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return max with multiple elements", func(t *testing.T) {
		n := []int{1, 3, -1, -3, 5, 3, 6, 7}
		k := 4
		expected := []int{3, 5, 5, 6, 7}
		actual := slidingwindowmax.SlidingWindowMax(n, k)

		assert.Equal(t, expected, actual)
	})
}
