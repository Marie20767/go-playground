package stacks_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Marie20767/go-playground/algorithms/arrays/stacks"
)

func TestIsStackValid(t *testing.T) {
	t.Run("should identify 1 valid bracket", func(t *testing.T) {
		s := "[]"
		actual := stacks.IsStackValid(s)
		expected := true

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify multiple different valid brackets", func(t *testing.T) {
		s := "([{}])"
		actual := stacks.IsStackValid(s)
		expected := true

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify multiple identical brackets", func(t *testing.T) {
		s := "{{{{{{}}}}}}"
		actual := stacks.IsStackValid(s)
		expected := true

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify invalid bracket sequence", func(t *testing.T) {
		s := "([)]"

		actual := stacks.IsStackValid(s)
		expected := false

		assert.Equal(t, expected, actual)
	})
}