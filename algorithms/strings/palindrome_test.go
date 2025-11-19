package strings_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	t.Run("should identify one palindrome", func(t *testing.T) {
		s := "a"
		expected := []string{"a"}
		actual := strings.PalindromeSubStrings(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify no palindrome", func(t *testing.T) {
		s := ""
		expected := []string{}
		actual := strings.PalindromeSubStrings(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify two palindromes", func(t *testing.T) {
		s := "ab"
		expected := []string{"a", "b"}
		actual := strings.PalindromeSubStrings(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify multiple palindromes", func(t *testing.T) {
		s := "abadd"
		expected := []string{"a", "aba", "b", "a", "d", "dd", "d"}
		actual := strings.PalindromeSubStrings(s)

		assert.Equal(t, expected, actual)
	})
}
