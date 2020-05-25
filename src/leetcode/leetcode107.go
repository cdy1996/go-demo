package main

import (
	"container/list"
	"fmt"
)

func main() {

}

// 树转二维数组
func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	// 使用双向队列
	list := list.New()
	// 头部插入
	list.PushFront(root)
	// 进行广度搜索
	for list.Len() > 0 {
		var currentLevel []int
		listLength := list.Len()
		for i := 0; i < listLength; i++ {
			// 尾部移除
			node := list.Remove(list.Back()).(*TreeNode)
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				list.PushFront(node.Left)
			}
			if node.Right != nil {
				list.PushFront(node.Right)
			}
		}
		result = append([][]int{currentLevel}, result...)

	}
	return result
}
