package flatten_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays/flatten"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	t.Run("should flatten 1 nested array", func(t *testing.T) {
		elements := []any{"hello", []any{"my", "name"}, "is", "something"}
		expected := []any{"hello", "my", "name", "is", "something"}
		actual := flatten.Flatten(elements)

		assert.Equal(t, expected, actual)
	})

	t.Run("should flatten 2 nested arrays", func(t *testing.T) {
		elements := []any{1, []any{2, []any{3, 4}, 5}}
		expected := []any{1, 2, 3, 4, 5}
		actual := flatten.Flatten(elements)

		assert.Equal(t, expected, actual)
	})

	t.Run("should flatten multiple nested arrays", func(t *testing.T) {
		elements := []any{1, []any{2, []any{3, 4, []any{5}}, 6}}
		expected := []any{1, 2, 3, 4, 5, 6}
		actual := flatten.Flatten(elements)

		assert.Equal(t, expected, actual)
	})

	t.Run("should flatten simple array", func(t *testing.T) {
		elements := []any{"hello"}
		expected := []any{"hello"}
		actual := flatten.Flatten(elements)

		assert.Equal(t, expected, actual)
	})
}
