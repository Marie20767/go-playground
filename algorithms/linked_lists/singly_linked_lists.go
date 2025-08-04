package singly_linked_lists

import "github.com/Marie20767/go-playground/data_structures/linked_list"

func ReverseList(node *linked_list.ListNode) *linked_list.ListNode {
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