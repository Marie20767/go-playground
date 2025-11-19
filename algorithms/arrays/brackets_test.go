package arrays_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays"
	"github.com/stretchr/testify/assert"
)

func TestAreBracketsValid(t *testing.T) {
	t.Run("should identify valid 1 valid bracket", func(t *testing.T) {
		s := "[]"
		actual := arrays.AreBracketsValid(s)
		expected := true

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify multiple different valid brackets", func(t *testing.T) {
		s := "([{}])"
		actual := arrays.AreBracketsValid(s)
		expected := true

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify multiple identical brackets", func(t *testing.T) {
		s := "{{{{{{}}}}}}"
		actual := arrays.AreBracketsValid(s)
		expected := true

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify invalid bracket sequence", func(t *testing.T) {
		s := "([)]"

		actual := arrays.AreBracketsValid(s)
		expected := false

		assert.Equal(t, expected, actual)
	})
}
