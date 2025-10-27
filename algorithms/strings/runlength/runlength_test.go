package runlength_test

import (
	"fmt"
	"testing"

	"github.com/Marie20767/go-playground/algorithms/strings/runlength"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input          string
	expectedOutput string
}

func TestRunLength(t *testing.T) {
	testCases := []testCase{
		{
			input:          "aabbcde",
			expectedOutput: "2a2b1c1d1e",
		},
		{
			input:          "wwwbbbw",
			expectedOutput: "3w3b1w",
		},
		{
			input:          "",
			expectedOutput: "",
		},
		{
			input:          "abcdefgh",
			expectedOutput: "1a1b1c1d1e1f1g1h",
		},
		{
			input:          "aaaaaaaa",
			expectedOutput: "8a",
		},
		{
			input: "aabbcdd",
			expectedOutput: "2a2b1c2d",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Should return %s for case %d", tc.expectedOutput, i), func(t *testing.T) {
			expected := tc.expectedOutput
			actual := runlength.Runlength(tc.input)

			assert.Equal(t, expected, actual)
		})
	}
}
