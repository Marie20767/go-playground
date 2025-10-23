package happynum_test

import (
	"testing"

	happynum "github.com/Marie20767/go-playground/algorithms/happynum"
	"github.com/stretchr/testify/assert"
)

func TestHappyNum(t *testing.T) {
	t.Run("Should identify happy num", func(t *testing.T) {
		n := 19
		expected := true
		actual := happynum.IsHappyNum(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should identify happy num below 10", func(t *testing.T) {
		n := 19
		expected := true
		actual := happynum.IsHappyNum(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should identify unhappy num below 10", func(t *testing.T) {
		n := 2
		expected := false
		actual := happynum.IsHappyNum(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should identify unhappy num above 10", func(t *testing.T) {
		n := 11
		expected := false
		actual := happynum.IsHappyNum(n)

		assert.Equal(t, expected, actual)
	})
}
