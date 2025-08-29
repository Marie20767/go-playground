package digitbalance_test

import (
	"testing"

	digitbalance "github.com/Marie20767/go-playground/algorithms/digit_balance"
	"github.com/stretchr/testify/assert"
)

func TestDigitBalance(t *testing.T) {
	t.Run("sum of first half equals sum of second half with even number of digits", func(t *testing.T) {
		n := 123303
		expected := true
		actual, err := digitbalance.DigitBalance(n)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("sum of first half equals sum of second half with odd number of digits", func(t *testing.T) {
		n := 1234321
		expected := true
		actual, err := digitbalance.DigitBalance(n)
		
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("sum of first half equals does not equal sum of second half", func(t *testing.T) {
		n := 1201100
		expected := false
		actual, err := digitbalance.DigitBalance(n)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
