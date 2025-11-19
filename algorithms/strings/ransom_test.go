package strings_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestRansom(t *testing.T) {
	t.Run("Should return true if notes and magazine are empty", func(t *testing.T) {
		note := ""
		magazine := ""
		expected := true
		actual := strings.Ransom(note, magazine)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return false with 1 missing letter in magazine", func(t *testing.T) {
		note := "hello"
		magazine := "ohelp"
		expected := false
		actual := strings.Ransom(note, magazine)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return true", func(t *testing.T) {
		note := "hello"
		magazine := "yellowh"
		expected := true
		actual := strings.Ransom(note, magazine)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should return true with long magazine", func(t *testing.T) {
		note := "hello"
		magazine := "sdyjfkdlsfjklswfjdksljdslfellowh"
		expected := true
		actual := strings.Ransom(note, magazine)

		assert.Equal(t, expected, actual)
	})
}
