package singlylinkedlists

import "github.com/Marie20767/go-playground/datastructures/linkedlist"

func ReverseList(node *linkedlist.ListNode) *linkedlist.ListNode {
	if node == nil {
		return nil
	}

	newHeadNode := node

	if node.Next != nil {
		newHeadNode = ReverseList(node.Next)
		node.Next.Next = node
	}

	node.Next = nil

	return newHeadNode
}
