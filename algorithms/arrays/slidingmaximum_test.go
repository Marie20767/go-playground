package arrays_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays"
	"github.com/stretchr/testify/assert"
)

func TestSlidingMax(t *testing.T) {
	t.Run("Should return empty array", func(t *testing.T) {
		n := []int{}
		expected := []int{}
		actual := arrays.SlidingMax(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return array with 1 element", func(t *testing.T) {
		n := []int{-100}
		expected := []int{-100}
		actual := arrays.SlidingMax(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return array with multiple elements", func(t *testing.T) {
		n := []int{1, 3, -1, -3, 5, 3, 6, 7}
		expected := []int{3, 3, -1, 5, 5, 6, 7}
		actual := arrays.SlidingMax(n)

		assert.Equal(t, expected, actual)
	})
}
