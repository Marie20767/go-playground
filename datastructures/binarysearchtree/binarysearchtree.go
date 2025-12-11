package binarysearchtree

import (
	"errors"
	"fmt"
)

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
	return b.Left == nil && b.Right == nil
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

// prints from smallest to largest
// go through left most part of the tree and print value, go back up to parent print value go to right and start process again
// example: 1 2 3 6 7 19
func (b *BST) DFSTraversal() {
	if b == nil {
		return
	}

	b.Left.DFSTraversal()
	fmt.Println(b.Value)
	b.Right.DFSTraversal()
}

// go through tree and check if you have an invalid value on left/right e.g. left bigger than value or right smaller than value
func (b *BST) IsValid() bool {
	if b == nil || b.IsLeafNode() {
		return true
	}

	if b.Left != nil && b.Value < b.Left.Value {
		return false
	}

	if b.Right != nil && b.Value > b.Right.Value {
		return false
	}

	if !b.Left.IsValid() {
		return false
	}

	return b.Right.IsValid()
}

// pass array and create tree from that array
func NewFromList(data []int) *BST {
	return newFromListHelper(data, 0, len(data)-1)
}

func newFromListHelper(data []int, start, end int) *BST {
	if start > end {
		return nil
	}

	midpoint := (end-start)/2 + start
	b := New(data[midpoint])
	b.Left = newFromListHelper(data, start, midpoint-1)
	b.Right = newFromListHelper(data, midpoint+1, end)

	return b
}

// TODO: implement
func (b *BST) Remove(value int) *BST {
	return b
}

// print level by level from left to right starting at the root
// example:
// 5 3 10 1 6 7
func (b *BST) BFSTraversal() {
	queue := []*BST{}
	queue = append(queue, b)

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		fmt.Println(first.Value)

		if first.Left != nil {
			queue = append(queue, first.Left)
		}

		if first.Right != nil {
			queue = append(queue, first.Right)
		}
	}
}