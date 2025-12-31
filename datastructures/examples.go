package datastructures

import (
	"fmt"

	"github.com/Marie20767/go-playground/datastructures/binarysearchtree"
	"github.com/Marie20767/go-playground/datastructures/graphs/matrix"
)

func BSTSearchExamples() {
	bst := binarysearchtree.NewFromList([]int{5, 3, 10, 1, 6, 7})
	fmt.Println(bst)
	bst.BFSTraversal()
}

func MatrixExamples() {
	data := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}
	m := matrix.New(data)
	fmt.Println(m)

	CanReachEndBFS := m.CanReachEnd2BFS()
	fmt.Println(">>> CanReachEndBFS: ", CanReachEndBFS)
}
