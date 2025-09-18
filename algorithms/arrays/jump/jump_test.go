package jump_test

import (
	"fmt"
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays/jump"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input          []int
	expectedOutput bool
}

func TestCanJump(t *testing.T) {

	testCases := []testCase{
		{
			input:          []int{3, 4, 0, 0, 0, 1, 0},
			expectedOutput: true,
		},
		{
			input:          []int{2, 5, 1, 0, 0, 0},
			expectedOutput: true,
		},
		{
			input:          []int{3, 4, 0, 0, 0, 1, 0, 0},
			expectedOutput: false,
		},
		{
			input:          []int{4, 3, 0, 2, 0, 1, 5},
			expectedOutput: true,
		},
		{
			input:          []int{0, 5, 5, 5},
			expectedOutput: false,
		},
		{
			input:          []int{1, 1, 1, 0},
			expectedOutput: true,
		},
		{
			input:          []int{1, 3, 2, 1},
			expectedOutput: true,
		},
		{
			input:          []int{5, 0, 0, 0, 0, 0, 9999},
			expectedOutput: false,
		},
		{
			input:          []int{},
			expectedOutput: true,
		},
		{
			input:          []int{2},
			expectedOutput: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Should return %t for case %d", tc.expectedOutput, i), func(t *testing.T) {
			expected := tc.expectedOutput
			actual := jump.CanJump(tc.input)

			assert.Equal(t, expected, actual)
		})
	}
}
