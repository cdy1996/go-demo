package main

func main() {

}
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		t1.Val = t1.Val + t2.Val
		t1.Left = mergeTrees(t1.Left, t2.Left)
		t1.Right = mergeTrees(t1.Right, t2.Right)
	} else if t1 == nil && t2 != nil {
		t1 = &TreeNode{
			Right: nil,
			Left:  nil,
			Val:   t2.Val,
		}
		t1.Left = mergeTrees(t1.Left, t2.Left)
		t1.Right = mergeTrees(t1.Right, t2.Right)
	} else if t1 != nil && t2 == nil {
		t1.Left = mergeTrees(t1.Left, nil)
		t1.Right = mergeTrees(t1.Right, nil)
	}
	return t1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
