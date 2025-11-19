package search_test

import (
	"fmt"
	"testing"

	"github.com/Marie20767/go-playground/algorithms/search"
	"github.com/Marie20767/go-playground/algorithms/search/helpers"
)

func TestBinarySearchRecursive(t *testing.T) {
	for label, testcase := range helpers.SearchTestCases {
		t.Run(fmt.Sprintf("binary search: %s", label), func(t *testing.T) {
			targetIndex, found := search.BinarySearchRecursive(testcase.Input, testcase.Target)
			if targetIndex != testcase.Expected {
				t.Fatalf("target index not expected: got %d, want %d", targetIndex, testcase.Expected)
			}
			if found != testcase.Found {
				t.Fatalf("found not equal: got %v, want %v", found, testcase.Found)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	for label, testcase := range helpers.SearchTestCases {
		t.Run(fmt.Sprintf("binary search: %s", label), func(t *testing.T) {
			targetIndex, found := search.BinarySearch(testcase.Input, testcase.Target)
			if targetIndex != testcase.Expected {
				t.Fatalf("target index not expected: got %d, want %d", targetIndex, testcase.Expected)
			}
			if found != testcase.Found {
				t.Fatalf("found not equal: got %v, want %v", found, testcase.Found)
			}
		})
	}
}
