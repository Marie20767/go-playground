package sorting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Marie20767/go-playground/algorithms/sorting"
)

func TestBubbleSort(t *testing.T) {
	t.Run("should sort numbers into ascending order", func(t *testing.T) {
		unsorted := []int{5, 1, 98, 17, 0, -100}
		expected := []int{-100, 0, 1, 5, 17, 98}

		actual := sorting.BubbleSort(unsorted)

		assert.Equal(t, expected, actual)
	})
}