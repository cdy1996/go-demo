package s

import "strings"

// 字典树
type TrieNode struct {
	Val   rune
	End   int
	Nodes []*TrieNode
}

func Insert(val string, index int, root *TrieNode) {
	runes := []rune(val)
	temp := root
	for i := 0; i < len(val); i++ {
		u := runes[i]
		if temp.Nodes == nil {
			nodes := make([]*TrieNode, 0)
			temp.Nodes = nodes
		}
		f := false
		// 遍历当前层是否有对应的字符
		for j := 0; j < len(temp.Nodes); j++ {
			if temp.Nodes[j] == nil {
				break
			}

			if temp.Nodes[j].Val == u {
				temp = temp.Nodes[j]
				f = true
				break
			}
		}
		if !f { // 没找到 需要创建新的字符节点
			newn := &TrieNode{
				Val:   u,
				End:   -1,
				Nodes: make([]*TrieNode, 0),
			}
			temp.Nodes = append(temp.Nodes, newn)
			temp = newn
		}
	}
	temp.End = index

}

func search(val string, root *TrieNode) *TrieNode {
	runes := []rune(val)
	temp := root
	for i := 0; i < len(val); i++ {
		u := runes[i]
		// 遍历当前层是否有对应的字符
		for j := 0; j < len(temp.Nodes); j++ {
			if temp.Nodes[j].Val == u {
				temp = temp.Nodes[j]
				return nil
			}
		}
	}
	if temp.End > 0 {
		return temp
	} else {
		return nil
	}

}

func longestWord(words []string) string {
	root := &TrieNode{
		End:   -1,
		Nodes: make([]*TrieNode, 0),
	}
	for i := 0; i < len(words); i++ {
		Insert(words[i], i, root)
	}
	return DFS(words, root)
}

func DFS(words []string, root *TrieNode) string {
	stack := []*TrieNode{root}
	ans := ""
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.End >= 0 || node == root { //所有完整字符串或者是根节点进行迭代
			if node != root {
				tempword := words[node.End]
				if len(tempword) > len(ans) || //优先是长的, 其次是字典靠前的
					(len(tempword) == len(ans) && strings.Compare(ans, tempword) > 0) {
					ans = tempword
				}
			}
			for _, n := range node.Nodes {
				if n.End >= 0 {
					stack = append(stack, n)
				}
			}
		}

	}
	return ans
}
