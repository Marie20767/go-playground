package arrays_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays"
	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	t.Run("Should find no anagrams", func(t *testing.T) {
		words := []string{"nine", "deed", "jo", "edee", "yes", "bye", "hello", "ibe", "heloh", "mine", "ok", "yessss"}
		expected := [][]string{}
		actual := arrays.Anagrams(words)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should find multiple anagrams", func(t *testing.T) {
		words := []string{"deed", "bye", "edde", "hello", "ibe", "mine", "ybe", "deed", "deedi"}
		expected := [][]string{{"deed", "edde", "deed"}, {"bye", "ybe"}}
		actual := arrays.Anagrams(words)

		assert.Equal(t, expected, actual)
	})
}
