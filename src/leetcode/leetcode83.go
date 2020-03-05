package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node5 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 3, Next: node5}
	node3 := &ListNode{Val: 2, Next: node4}
	node2 := &ListNode{Val: 1, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	head := deleteDuplicates(node1)
	for head != nil {
		fmt.Printf("%v \n", head.Val)
		head = head.Next
	}

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	a := head
	if a == nil {
		return head
	}
	b := a.Next
	if b == nil {
		return head
	}
	for b != nil {
		if a.Val == b.Val {
			b = b.Next
			a.Next = b
			if b == nil {
				break
			}
		} else {
			a = a.Next
			b = b.Next
		}

	}
	return head
}
