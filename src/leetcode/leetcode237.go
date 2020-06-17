package main

/**
删除链表中的节点
*/
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	if next.Next == nil {
		node.Next = nil
		return
	} else {
		node.Next = next.Next
	}
}
