package sort_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/Marie20767/go-playground/algorithms/sort"
	"github.com/Marie20767/go-playground/algorithms/sort/helpers"
)

func TestQuickSortNotInPlace(t *testing.T) {
	for label, testcase := range helpers.SortingTestCases {
		t.Run(fmt.Sprintf("quick sort: %s", label), func(t *testing.T) {
			// Copy the input because the test data is shared among sorting tests
			inputCopy := make([]int, len(testcase.Input))
			copy(inputCopy, testcase.Input)

			actual := sort.QuickSortNotInPlace(inputCopy)
			if !slices.Equal(actual, testcase.Expected) {
				t.Fatalf("slices not equal: got %v, want %v", actual, testcase.Expected)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for label, testcase := range helpers.SortingTestCases {
		t.Run(fmt.Sprintf("quick sort: %s", label), func(t *testing.T) {
			// Copy the input because the test data is shared among sorting tests
			inputCopy := make([]int, len(testcase.Input))
			copy(inputCopy, testcase.Input)

			sort.QuickSort(inputCopy)
			if !slices.Equal(inputCopy, testcase.Expected) {
				t.Fatalf("slices not equal: got %v, want %v", inputCopy, testcase.Expected)
			}
		})
	}
}
