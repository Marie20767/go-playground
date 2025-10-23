package alphabetstepper_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings/alphabetstepper"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	t.Run("should shift every letter in lowercase string to next letter", func(t *testing.T) {
		s := "abcdzhj"
		expected := "bcdeaik"
		actual := alphabetstepper.AlphabetStepper(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should shift every letter in mixed cased string to next letter", func(t *testing.T) {
		s := "abcDzhJ"
		expected := "bcdEaiK"
		actual := alphabetstepper.AlphabetStepper(s)

		assert.Equal(t, expected, actual)
	})
}
