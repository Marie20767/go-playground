package datastructures

import (
	"fmt"

	"github.com/Marie20767/go-playground/datastructures/binarysearchtree"
)

func BSTSearchExample() {
	bst := binarysearchtree.NewFromList([]int{1, 3, 6, 7, 19})

	fmt.Println(bst)

	bWithInsertion := bst.Insert(18)
	fmt.Println(bWithInsertion)
}
