package main

import (
	"fmt"
)

func main() {

	node4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}

	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	swapPairs(node1)
	fmt.Print(node1.Val)

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	next := head.Next
	if next == nil {
		return head
	} else {
		head.Next = swapPairs(next.Next)
	}
	next.Next = head
	return next

}
