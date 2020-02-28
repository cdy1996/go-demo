package main

import (
	"fmt"
)

func main() {
	//fmt.Printf("%v", reflect.ValueOf("str"[1]))
	//fmt.Printf("%v", reflect.TypeOf("str"[1]))
	fmt.Printf("%v", longestCommonPrefix([]string{"aa", "a"}))
	fmt.Printf("%v", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("%v", longestCommonPrefix([]string{}))
	fmt.Printf("%v", longestCommonPrefix([]string{"dog"}))
	fmt.Printf("%v", longestCommonPrefix([]string{"dog", "dog"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	j := 0
I:
	for {
		if j > len(strs[0])-1 {
			break I
		}
		var c = strs[0][j]
		for i := 1; i < len(strs); i++ {
			// 字符已到最后 或者 字符与当期字符的j位置不相等
			if j > len(strs[i])-1 || c != strs[i][j] {
				break I
			}
			if i == len(strs)-1 { //比较下一位
				j++
			}
		}
	}
	if j == 0 {
		return ""
	}
	return strs[0][0:j]
}
