package main

/**
返回倒数第 k 个节点
https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
*/
func main() {

}

func kthToLast(head *ListNode, k int) int {

	slow := head
	fast := head
	for i := 1; i < k; i++ {
		fast = fast.Next
	}
	for {
		if fast.Next == nil {
			return slow.Val
		}
		slow = slow.Next
		fast = fast.Next
	}

}
