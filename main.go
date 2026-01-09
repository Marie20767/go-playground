package main

import (
	"github.com/Marie20767/go-playground/algorithms/graphs/treasure"
)

func main() {
	// datastructures.BSTSearchExamples()
	// datastructures.MatrixExamples()
	// concurrency.ConcurrencyExamples()

	treasure.IslandsAndTreasure([][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	})
}
