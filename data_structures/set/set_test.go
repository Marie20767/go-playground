package set_test

import (
	"testing"

	"github.com/Marie20767/go-playground/data_structures/set"
	"github.com/stretchr/testify/assert"
)

func TestSet (t *testing.T) {
	t.Run("should add unique numbers to empty set", func(t *testing.T) {
		newSet := set.NewSet()
		newSet.Add(1, 2, 3, 2, 1)

		expected := []int{1, 2, 3}
		actual := newSet.Data

		assert.Equal(t, expected, actual)
	})

	t.Run("should add unique numbers to set", func(t *testing.T) {
		newSet := set.NewSet(5, 1, 6)
		newSet.Add(1, 2, 3, 2, 1)

		expected := []int{5, 1, 6, 2, 3}
		actual := newSet.Data

		assert.Equal(t, expected, actual)
	})

	t.Run("should remove number from set", func(t *testing.T) {
		newSet := set.NewSet(5, 1, 6, 0)
		newSet.Delete(6)

		expected := []int{5, 1, 0}
		actual := newSet.Data

		assert.Equal(t, expected, actual)
	})
}