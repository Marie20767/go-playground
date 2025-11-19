package arrays_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays"
	"github.com/stretchr/testify/assert"
)

func TestPeakFinder(t *testing.T) {
	t.Run("should return 1 peak", func(t *testing.T) {
		n := []int{0, 1}
		expected := []int{1}
		actual := arrays.PeakFinder(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return multiple peaks", func(t *testing.T) {
		n := []int{0, 1, 0, 2, 3, 1, 4}
		expected := []int{1, 3, 4}
		actual := arrays.PeakFinder(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return multiple peaks with long nums list", func(t *testing.T) {
		n := []int{10, 1, 5, 100, 1, 1, 8, 999, 0, 55, 7}
		expected := []int{10, 100, 999, 55}
		actual := arrays.PeakFinder(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return multiple peaks with negative num", func(t *testing.T) {
		n := []int{-10, 1, 5, 100, 1, 1, 8, 999, 0, 55, 7}
		expected := []int{100, 999, 55}
		actual := arrays.PeakFinder(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return empty peak", func(t *testing.T) {
		n := []int{0}
		expected := []int{}
		actual := arrays.PeakFinder(n)

		assert.Equal(t, expected, actual)
	})

}
