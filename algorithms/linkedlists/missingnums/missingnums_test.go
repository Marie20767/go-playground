package missingnums_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/linkedlists/missingnums"
	"github.com/stretchr/testify/assert"
)

func TestMissingNum(t *testing.T) {
	t.Run("Should return missing number 0", func(t *testing.T) {
		n := []int{}
		expected := 0
		actual := missingnums.MissingNums(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return missing number 1", func(t *testing.T) {
		n := []int{0}
		expected := 1
		actual := missingnums.MissingNums(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return missing number 8", func(t *testing.T) {
		n := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
		expected := 8
		actual := missingnums.MissingNums(n)

		assert.Equal(t, expected, actual)
	})
}
