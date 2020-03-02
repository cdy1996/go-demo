package main

import (
	"fmt"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l11 := l1
	l22 := l2
	var head = &ListNode{}
	h := head
	for l11 != nil && l22 != nil {

		if l11.Val <= l22.Val {
			l := &ListNode{
				Val: l11.Val,
			}
			head.Next = l
			head = l
			l11 = l11.Next
		} else {
			l := &ListNode{
				Val: l22.Val,
			}
			head.Next = l
			head = l
			l22 = l22.Next
		}
	}
	if l11 == nil {
		head.Next = l22
	}
	if l22 == nil {
		head.Next = l11
	}
	return h.Next
}
