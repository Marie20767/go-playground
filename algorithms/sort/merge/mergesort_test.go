package mergesort_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/Marie20767/go-playground/algorithms/sort/helpers"
	"github.com/Marie20767/go-playground/algorithms/sort/merge"
)

func TestMergeSort(t *testing.T) {
	for label, testcase := range helpers.SortingTestCases {
		t.Run(fmt.Sprintf("merge sort: %s", label), func(t *testing.T) {
			// Copy the input because the test data is shared among sorting tests
			inputCopy := make([]int, len(testcase.Input))
			copy(inputCopy, testcase.Input)

			actual := mergesort.MergeSort(inputCopy)
			if !slices.Equal(actual, testcase.Expected) {
				t.Fatalf("slices not equal: got %v, want %v", actual, testcase.Expected)
			}
		})
	}
}
