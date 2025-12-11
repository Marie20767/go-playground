package datastructures

import (
	"fmt"

	"github.com/Marie20767/go-playground/datastructures/binarysearchtree"
)

func BSTSearchExample() {
	bst := binarysearchtree.NewFromList([]int{5, 3, 10, 1, 6, 7})

	fmt.Println(bst)

	bst.BFSTraversal()
}
