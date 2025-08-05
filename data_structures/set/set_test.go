package set_test

import (
	"testing"

	"github.com/Marie20767/go-playground/data_structures/set"
	"github.com/stretchr/testify/assert"
)

func TestSet (t *testing.T) {
	t.Run("should add unique numbers to empty set", func(t *testing.T) {
		newSet := set.NewSet()
		newSet.Add(1, 2, 3, 2, 1)

		expected := map[int]struct{}{ 1: {}, 2: {}, 3: {}}
		actual := newSet.Data

		assert.Equal(t, expected, actual)
	})

	t.Run("should add unique numbers to set", func(t *testing.T) {
		newSet := set.NewSet(5, 1, 6)
		newSet.Add(1, 2, 3, 2, 1)

		expected := map[int]struct{}{ 5: {}, 1: {}, 6: {}, 2: {}, 3: {}}
		actual := newSet.Data

		assert.Equal(t, expected, actual)
	})

	t.Run("should remove number from set", func(t *testing.T) {
		newSet := set.NewSet(5, 1, 6, 0)
		newSet.Delete(6)

		expected := map[int]struct{}{ 5: {}, 1: {}, 0: {}}
		actual := newSet.Data

		assert.Equal(t, expected, actual)
	})

	t.Run("should contain number", func(t *testing.T) {
		newSet := set.NewSet(5, 1, 6, 0)
		actual := newSet.Contains(6)
		expected := true

		assert.Equal(t, expected, actual)
	})	

	t.Run("should not contain number", func(t *testing.T) {
		newSet := set.NewSet(5, 1, 6, 0)
		actual := newSet.Contains(7)
		expected := false

		assert.Equal(t, expected, actual)
	})

	t.Run("should intersect 2 sets", func(t *testing.T) {
		setA := set.NewSet(5, 1, 6, 0)
		setB := set.NewSet(4, 1, 6, 10)

		actual := setA.Intersection(setB).Data
		expected := map[int]struct{}{ 1: {}, 6: {}}

		assert.Equal(t, expected, actual)
	})

	t.Run("should intersect 2 sets with 1 empty set", func(t *testing.T) {
		setA := set.NewSet()
		setB := set.NewSet(4, 1, 6, 10)

		actual := setA.Intersection(setB).Data
		expected := map[int]struct{}{}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return elements only unique to first set", func(t *testing.T) {
		setA := set.NewSet(1, 10, 5, 4)
		setB := set.NewSet(4, 1, 6, 10)

		actual := setA.Difference(setB).Data
		expected := map[int]struct{}{ 5: {}}

		assert.Equal(t, expected, actual)
	})

	t.Run("should return all elements from the first set", func(t *testing.T) {
		setA := set.NewSet(1, 10, 5, 4)
		setB := set.NewSet()

		actual := setA.Difference(setB).Data
		expected := map[int]struct{}{ 1: {}, 10: {}, 5: {}, 4: {}}

		assert.Equal(t, expected, actual)
	})

	t.Run("should combine elements from both sets", func(t *testing.T) {
		setA := set.NewSet(1, 10, 5, 4)
		setB := set.NewSet(1, 6, 5)

		actual := setA.Union(setB).Data
		expected := map[int]struct{}{ 1: {}, 10: {}, 5: {}, 6: {}, 4: {}}

		assert.Equal(t, expected, actual)
	})

	t.Run("should combine elements from both sets with 1 empty set", func(t *testing.T) {
		setA := set.NewSet()
		setB := set.NewSet(1, 6, 5)

		actual := setA.Union(setB).Data
		expected := map[int]struct{}{ 1: {}, 6: {}, 5: {}}

		assert.Equal(t, expected, actual)
	})

	t.Run("should list nums from set", func(t *testing.T) {
		newSet := set.NewSet(1, 5, 5, 10, 8, 9)

		actual := newSet.ToList()
		expected := []int{1, 5, 10, 8, 9}

		assert.Equal(t, expected, actual)
	})

	t.Run("should list nums from empty set", func(t *testing.T) {
		newSet := set.NewSet()

		actual := newSet.ToList()
		expected := []int{}

		assert.Equal(t, expected, actual)
	})


	t.Run("should return nums as formatted string", func(t *testing.T) {
		newSet := set.NewSet(5, 10, 110)

		actual := newSet.String()
		expected := "5, 10, 110"

		assert.Equal(t, expected, actual)
	})

	t.Run("should return empty string", func(t *testing.T) {
		newSet := set.NewSet()

		actual := newSet.String()
		expected := ""

		assert.Equal(t, expected, actual)
	})
}