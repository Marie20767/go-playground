package movezeroes_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays/movezeroes"
	"github.com/stretchr/testify/assert"
)

func TestRemoveZeroes(t *testing.T) {
	t.Run("Should move all zeroes to the end", func(t *testing.T) {
		n := []int{0, 0, 1, 10, 0, 0, 50}
		expected := []int{1, 10, 50, 0, 0, 0, 0}
		movezeroes.MoveZeroesInPlace(n)

		assert.Equal(t, expected, n)
	})

	t.Run("Should return empty slice", func(t *testing.T) {
		n := []int{}

		expected := []int{}
		movezeroes.MoveZeroesInPlace(n)

		assert.Equal(t, expected, n)
	})
}
