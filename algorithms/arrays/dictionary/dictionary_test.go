package dictionary_test

import (
	"fmt"
	"testing"

	"github.com/Marie20767/go-playground/algorithms/arrays/dictionary"
	"github.com/stretchr/testify/assert"
)

type testCases struct {
	input          []string
	expectedOutput int
}

func TestCharRemoval(t *testing.T) {
	testCases := []testCases{
		{
			input:          []string{"baseball", "a,all,b,ball,bas,base,cat,code,d,e,quit,z"},
			expectedOutput: 4,
		},
		{
			input:          []string{"apbpleeeef", "a,ab,abc,abcg,b,c,dog,e,efd,zzzz"},
			expectedOutput: 8,
		},
		{
			input:          []string{"apple", "cot,c,cod,t"},
			expectedOutput: -1,
		},
		{
			input:          []string{"", ""},
			expectedOutput: -1,
		},
		{
			input: []string{"wworlcde", "apple,bat,cat,goodbye,hello,yellow,why,world"},
			expectedOutput: 3,
		},
		{
			input: []string{"wwerlcdo", "apple,bat,cat,goodbye,hello,yellow,why,world"},
			expectedOutput: -1,
		},
		{
			input: []string{"werld", "a"},
			expectedOutput: -1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Should return %d for case %d", tc.expectedOutput, i), func(t *testing.T) {
			expected := tc.expectedOutput
			actual := dictionary.CharRemoval(tc.input)

			assert.Equal(t, expected, actual)
		})
	}
}
