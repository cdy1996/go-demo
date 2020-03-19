package main

import (
	"fmt"
	"rpc/src/leetcode/s"
	"strings"
)

func main() {
	//fmt.Printf("%v \n", longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	//fmt.Printf("%v \n", longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
	fmt.Printf("%v \n", longestWord([]string{"t", "ti", "tig", "tige", "tiger", "e", "en", "eng", "engl", "engli", "englis", "english", "h", "hi", "his", "hist", "histo", "histor", "history"}))

}

// 最长的可遍历单词， 构建字典树
func longestWord(words []string) string {
	root := &s.TrieNode{
		End:   -1,
		Nodes: make([]*s.TrieNode, 0),
	}
	for i := 0; i < len(words); i++ {
		s.Insert(words[i], i, root)
	}
	return DFS(words, root)
}

func DFS(words []string, root *s.TrieNode) string {
	stack := []*s.TrieNode{root}
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
