package distinct_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings/distinct"
	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	t.Run("should identify distinct substring with 1 letter", func(t *testing.T) {
		s := "aaaaaa"
		expected := "a"
		actual := distinct.LargestDistinct(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify no distinct substring", func(t *testing.T) {
		s := ""
		expected := ""
		actual := distinct.LargestDistinct(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify longest distinct substring", func(t *testing.T) {
		s := "axayazba"
		expected := "yazb"
		actual := distinct.LargestDistinct(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify longest distinct substring with distinct string", func(t *testing.T) {
		s := "abcd"
		expected := "abcd"
		actual := distinct.LargestDistinct(s)

		assert.Equal(t, expected, actual)
	})
}
