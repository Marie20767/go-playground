package filter_test

import (
	"testing"

	"github.com/Marie20767/go-playground/higher_order/filter"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("should filter out any nums <= 2", func(t *testing.T) {
		elements := []int{0, 1, 10, 0, 2, 2, 1, 500}
		predicate := func(i int) bool {
			return i > 2
		}
		expected := []int{10, 500}
		actual := filter.Filter(elements, predicate)

		assert.Equal(t, expected, actual)
	})


	t.Run("should filter out uneven nums", func(t *testing.T) {
		elements := []int{21, 0, 1, 10, 0, 2, 2, 1, 500, 4}
		predicate := func(i int) bool {
			return i % 2 == 0
		}
		expected := []int{0, 10, 0, 2, 2, 500, 4}
		actual := filter.Filter(elements, predicate)

		assert.Equal(t, expected, actual)
	})

	t.Run("should filter out all nums", func(t *testing.T) {
		elements := []int{21, 0, 1, 10, 0, 2, 2, 1, 500, 4}
		predicate := func(i int) bool {
			return i > 1000
		}
		expected := []int{}
		actual := filter.Filter(elements, predicate)

		assert.Equal(t, expected, actual)
	})

	t.Run("should filter out strings longer than 5 chars", func(t *testing.T) {
		elements := []string{"hello", "something", "bye", "now", "please!"}
		predicate := func(s string) bool {
			return len(s) <= 5
		}
		expected := []string{"hello", "bye", "now"}
		actual := filter.Filter(elements, predicate)

		assert.Equal(t, expected, actual)
	})
}
