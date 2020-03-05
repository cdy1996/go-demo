package main

import (
	"fmt"
)

func main() {

	node2 := &TreeNode{
		Val: 2,
	}
	node22 := &TreeNode{
		Val: 3,
	}
	node33 := &TreeNode{
		Val: 3,
	}
	node3 := &TreeNode{
		Val: 2,
	}
	node := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node33,
	}
	node1 := &TreeNode{
		Val:   1,
		Left:  node3,
		Right: node22,
	}
	fmt.Printf("%v \n", isSameTree(node, node1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Right)

}
