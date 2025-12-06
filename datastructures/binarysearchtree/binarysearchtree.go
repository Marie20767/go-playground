package binarysearchtree

import "errors"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func New(value int) *BST {
	return &BST{
		Value: value,
	}
}

func (b *BST) SetLeft(v int) error {
	if v > b.Value {
		return errors.New("left value has to be smaller than root value")
	}

	b.Left = New(v)

	return nil
}

func (b *BST) SetRight(v int) error {
	if v < b.Value {
		return errors.New("right value has to be larger than root value")
	}

	b.Right = New(v)

	return nil
}

func (b *BST) IsLeafNode() bool {
	return b.Left != nil || b.Right != nil
}

func (b *BST) FindMin() *BST {
	if b.Left == nil {
		return b
	}

	return b.Left.FindMin()
}

func (b *BST) FindMax() *BST {
	if b.Right == nil {
		return b
	}

	return b.Right.FindMax()
}

func (b *BST) Find(value int) *BST {
	if b == nil {
		return nil
	}

	if b.Value == value {
		return b
	}

	if value < b.Value {
		return b.Left.Find(value)
	}

	return b.Right.Find(value)
}

func (b *BST) Insert(value int) *BST {
	if b == nil {
		return New(value)
	}

	if value < b.Value {
		b.Left = b.Left.Insert(value)
	} else {
		b.Right = b.Right.Insert(value)
	}

	return b
}

func (b *BST) DFSTraversal(f func(value int)) {
	// prints from smallest to largest
	// example:
	// 1 3 5 6 7 10

	// go as low as you can on the left
	// print entire left subtree
	// print parent
	// check right
	// print entire left subtree
	// etc
}

func (b *BST) IsValid() bool {
	// go through tree and check if you have an invalid value on left/right e.g. left bigger than value or right smaller than value
	return true
}

// pass array and create tree from that array
func NewFromList(data []int) *BST {
	return newFromListHelper(data, 0, len(data)-1)
}

func newFromListHelper(data []int, start, end int) *BST {
	if start > end {
		return nil
	}

	midpoint := (end - start) / 2 + start
	b := New(data[midpoint])
	b.Left = newFromListHelper(data, start, midpoint-1)
	b.Right = newFromListHelper(data, midpoint+1, end)

	return b
}

func (b *BST) Remove(value int) *BST {}

func (b *BST) BFSTraversal(f func(value int)) {
	// print level by level from left to right starting at the root
	// example:
	// 5 3 10 1 6 7
}
