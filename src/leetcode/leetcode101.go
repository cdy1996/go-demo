package main

import (
	"fmt"
)

func main() {

	node3 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node1 := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	fmt.Printf("%v \n", isSymmetric(node1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	return isSameNode(left, right)
}

func isSameNode(p, q *TreeNode) bool {
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	return (p == nil && q == nil) || (p.Val == q.Val && isSameNode(p.Left, q.Right) && isSameNode(p.Right, q.Left))
}
