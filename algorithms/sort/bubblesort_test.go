package sort_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/Marie20767/go-playground/algorithms/sort"
	"github.com/Marie20767/go-playground/algorithms/sort/helpers"
)

func TestBubbleSort(t *testing.T) {
	for label, testcase := range helpers.SortingTestCases {
		t.Run(fmt.Sprintf("bubble sort: %s", label), func(t *testing.T) {
			// Copy the input because the test data is shared among sorting tests
			inputCopy := make([]int, len(testcase.Input))
			copy(inputCopy, testcase.Input)

			sort.BubbleSort(inputCopy)
			if !slices.Equal(inputCopy, testcase.Expected) {
				t.Fatalf("slices not equal: got %v, want %v", inputCopy, testcase.Expected)
			}
		})
	}
}
