package main

import "fmt"

func main() {

	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	fmt.Printf("%v \n", hasCycle(l1))

}

//
//func hasCycle(head *ListNode) bool {
//	if head == nil {
//		return false
//	}
//	slow := head
//	if head.Next == nil {
//		return false
//
//	}
//	fast := head.Next
//
//	for slow!=fast{
//		if fast ==nil || fast.Next ==nil {
//			return false
//		}
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return true
//}

// 链表循环判断
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	if head.Next == nil {
		return false

	}
	fast := head.Next

	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next

		if fast.Next == nil {
			return false
		}
		if fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return false
}
