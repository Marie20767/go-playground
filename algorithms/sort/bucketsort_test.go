package sort_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/Marie20767/go-playground/algorithms/sort"
	"github.com/Marie20767/go-playground/algorithms/sort/helpers"
)

func TestBucketSort(t *testing.T) {
	for label, testcase := range helpers.BucketSortingTestCases {
		t.Run(fmt.Sprintf("bucket sort: %s", label), func(t *testing.T) {
			// Copy the input because the test data is shared among sorting tests
			inputCopy := make([]int, len(testcase.Input))
			copy(inputCopy, testcase.Input)

			sort.BucketSort(inputCopy)
			if !slices.Equal(inputCopy, testcase.Expected) {
				t.Fatalf("slices not equal: got %v, want %v", inputCopy, testcase.Expected)
			}
		})
	}
}

func TestBucketSortInRange(t *testing.T) {
	for label, testcase := range helpers.BucketSortingInRangeTestCases {
		t.Run(fmt.Sprintf("bucket sort in range: %s", label), func(t *testing.T) {
			// Copy the input because the test data is shared among sorting tests
			inputCopy := make([]int, len(testcase.Input))
			copy(inputCopy, testcase.Input)

			sort.BucketSortInRange(inputCopy)
			if !slices.Equal(inputCopy, testcase.Expected) {
				t.Fatalf("slices not equal: got %v, want %v", inputCopy, testcase.Expected)
			}
		})
	}
}
