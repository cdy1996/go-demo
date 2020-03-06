package main

import (
	"fmt"
)

func main() {

	node7 := &TreeNode{
		Val: 7,
	}
	node15 := &TreeNode{
		Val: 15,
	}
	node20 := &TreeNode{
		Val:   20,
		Left:  node15,
		Right: node7,
	}
	node9 := &TreeNode{
		Val: 9,
	}
	node3 := &TreeNode{
		Val:   3,
		Left:  node9,
		Right: node20,
	}
	fmt.Printf("%v \n", maxDepth(node3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	maxLeft := 0
	if root.Left != nil {
		maxLeft = 1 + maxDepth(root.Left)

	}
	maxRight := 0
	if root.Right != nil {
		maxRight = 1 + maxDepth(root.Right)
	}

	if maxLeft > maxRight {
		return maxLeft
	} else {
		return maxRight
	}
}

func maxDepthFor(root *TreeNode) int {
	var queue []*TreeNode

	queue = append(queue, root)
	level := 0 //层数

	preCount := 1 //前一层的数量
	pCount := 0   //下一层的数量
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1 : len(queue)-1]
		preCount--
		if node.Left != nil {
			queue = append(queue, node.Left)
			pCount++
		}
		if node.Right != nil {
			queue = append(queue, node.Left)
			pCount++
		}
		if preCount == 0 {
			preCount = pCount
			pCount = 0
			level++
		}
	}
	return level
}
