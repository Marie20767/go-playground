package reverse_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings/reverse"
	"github.com/stretchr/testify/assert"
)

func TestRevers(t *testing.T) {
	t.Run("should reverse string", func(t *testing.T) {
		s := "the sky is blue"
		expected := "blue is sky the"
		actual := reverse.Reverse(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should reverse string with whitespaces", func(t *testing.T) {
		s := "  hello     world  "
		expected := "world hello"
		actual := reverse.Reverse(s)

		assert.Equal(t, expected, actual)
	})
}