package dynamic_arrays_test

import (
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays"
	"github.com/stretchr/testify/assert"
)

func TestDynamicArrays(t *testing.T) {
	t.Run("should get element by index", func(t *testing.T) {
		newArray := dynamic_arrays.NewDynamicArray(3)

		newArray.Pushback(1)
		newArray.Pushback(2)
		newArray.Pushback(3)

		expected := 2
		actual := newArray.Get(1)

		assert.Equal(t, expected, actual)
	})

	t.Run("should set element at index", func(t *testing.T) {
		newArray := dynamic_arrays.NewDynamicArray(3)

		newArray.Pushback(0)
		newArray.Pushback(2)
		newArray.Pushback(3)

		newArray.Set(0, 1)

		expected := 1
		actual := newArray.Get(0)

		assert.Equal(t, expected, actual)
	})

	t.Run("should get array capacity", func(t *testing.T) {
		capacity := 1
		newArray := dynamic_arrays.NewDynamicArray(capacity)

		expected := capacity
		actual := newArray.GetCapacity()

		assert.Equal(t, expected, actual)
	})

	t.Run("should resize array", func(t *testing.T) {
		newArray := dynamic_arrays.NewDynamicArray(1)

		newArray.Pushback(1)
		newArray.Pushback(2)
		newArray.Pushback(3)

		expected := 4
		actual := newArray.GetCapacity()

		assert.Equal(t, expected, actual)

	})
}
