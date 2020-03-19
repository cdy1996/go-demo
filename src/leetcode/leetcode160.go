package main

import "fmt"

func main() {
	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l4,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l2,
	}
	l9 := &ListNode{
		Val:  9,
		Next: l1,
	}
	l0 := &ListNode{
		Val:  0,
		Next: l9,
	}

	fmt.Printf("%v", getIntersectionNode(l3, l0))

}

// 查找链表的交叉
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var tailA *ListNode
	var tailB *ListNode
	tempA := headA
	flagA := true //tempA是否要交替到B的链表
	tempB := headB
	flagB := true
	for {
		if tempA == tempB { // 遍历到同一个位置了
			return tempA
		}
		if tempA.Next == nil { //tempA 遍历到尾部了
			if tailA == nil {
				tailA = tempA
			}
			if flagA {
				tempA = headB // tempA切换到链表B的head
			} else {
				tempA = headA
			}
			flagA = !flagA
		} else {
			tempA = tempA.Next
		}
		if tempB.Next == nil {
			tailB = tempB
			if tailB == nil {
				tailB = tempB
			}
			if flagB {
				tempB = headA
			} else {
				tempB = headB
			}
			flagB = !flagB
		} else {
			tempB = tempB.Next
		}
		// 尾部相同时才会有交点
		if tailA != nil && tailB != nil && tailA != tailB {
			return nil
		}
	}

}
