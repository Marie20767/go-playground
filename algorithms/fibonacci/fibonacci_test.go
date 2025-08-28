package fibonacci_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/fibonacci"
	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	t.Run("should return 0 if n is 0", func(t *testing.T) {
		expected := 0
		n := 0
		actual := fibonacci.Fibonnaci(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct number in fibonacci sequence", func(t *testing.T) {
		expected := 34
		n := 9
		actual := fibonacci.Fibonnaci(n)

		assert.Equal(t, expected, actual)
	})
}

func TestFibRec(t *testing.T) {
	t.Run("should return 1 if n is 1", func(t *testing.T) {
		expected := 1
		n := 1
		actual := fibonacci.RecursiveFib(n)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return correct number in fibonacci sequence", func(t *testing.T) {
		expected := 55
		n := 10
		actual := fibonacci.RecursiveFib(n)

		assert.Equal(t, expected, actual)
	})
}
