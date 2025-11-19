package arrays_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays"
	"github.com/stretchr/testify/assert"
)

func TestNestedSum(t *testing.T) {
	t.Run("Should return sum with 1 nested slice", func(t *testing.T) {
		n := []any{1, 2, []any{3, 4}, 5}
		expected := 15
		actual := arrays.NestedSum(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return sum with 2 nested slices", func(t *testing.T) {
		n := []any{1, []any{2, []any{3, 4}, 5}}
		expected := 15
		actual := arrays.NestedSum(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return sum with multipe nested slices", func(t *testing.T) {
		n := []any{1, []any{2, []any{3, 4}, 5}}
		expected := 15
		actual := arrays.NestedSum(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return sum with empty slices", func(t *testing.T) {
		n := []any{[]any{1, 2}, []any{3, []any{4}}, 5}
		expected := 15
		actual := arrays.NestedSum(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return sum with empty slices", func(t *testing.T) {
		n := []any{[]any{}, 1, []any{2, []any{}}, 3}
		expected := 6
		actual := arrays.NestedSum(n)

		assert.Equal(t, expected, actual)
	})
}
