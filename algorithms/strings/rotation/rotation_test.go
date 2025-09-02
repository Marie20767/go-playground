package rotation_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings/rotation"
	"github.com/stretchr/testify/assert"
)

func TestIsRotation(t *testing.T) {
	t.Run("should identify double rotation", func(t *testing.T) {
		s1, s2 := "rotation", "tationro"
		expected := true
		actual := rotation.IsRotation(s1, s2)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify triple rotation", func(t *testing.T) {
		s1, s2 := "hello", "llohe"
		expected := true
		actual := rotation.IsRotation(s1, s2)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify rotation with empty strings", func(t *testing.T) {
		s1, s2 := "", ""
		expected := true
		actual := rotation.IsRotation(s1, s2)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify rotation with 1 character", func(t *testing.T) {
		s1, s2 := "a", "a"
		expected := true
		actual := rotation.IsRotation(s1, s2)

		assert.Equal(t, expected, actual)
	})

	t.Run("should identify incorrect rotation", func(t *testing.T) {
		s1, s2 := "abcde", "abced"
		expected := false
		actual := rotation.IsRotation(s1, s2)

		assert.Equal(t, expected, actual)
	})
}
