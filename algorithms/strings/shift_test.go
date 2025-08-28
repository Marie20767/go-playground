package shift_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	t.Run("should shift every letter in string to next letter", func(t *testing.T) {
		s := "abcdzhj"
		expected := "bcdeaik"
		actual := shift.Shift(s)

		assert.Equal(t, expected, actual)
	})
}
