package strings_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	t.Run("should return first unique character index", func(t *testing.T) {
		s := "loveleetcode"
		expected := 2
		actual := strings.FirstUnique(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return first unique emoji index", func(t *testing.T) {
		s := "🙂🙃🙂"
		expected := 1
		actual := strings.FirstUnique(s)

		assert.Equal(t, expected, actual)
	})

	t.Run("should find no unique character index", func(t *testing.T) {
		s := "aabb"
		expected := -1
		actual := strings.FirstUnique(s)

		assert.Equal(t, expected, actual)
	})
}
