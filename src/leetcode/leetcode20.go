package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v", isValid(""))
	fmt.Printf("%v", isValid("(("))
	fmt.Printf("%v", isValid("{[]}"))
	fmt.Printf("%v", isValid("{[][]}"))

}
func isValid(s string) bool {
	//maps := make(map[uint8]bool)
	//maps['('] = true
	//maps['['] = true
	//maps['{'] = true
	//maps[')'] = false
	//maps[']'] = false
	//maps['}'] = false
	var stack []uint8
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 { //奇数个
		return false
	}
	for i := 0; i < len(s); i++ {
		r := s[i] == '{' || s[i] == '(' || s[i] == '['
		b := s[i] == '}' || s[i] == ')' || s[i] == ']'
		if r || b {
			if r { // 括号
				stack = append(stack, s[i])
			} else { // 括回
				if len(stack) == 0 {
					return false
				}
				u := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				if (u == '(' && s[i] == ')') || (u == '{' && s[i] == '}') || (u == '[' && s[i] == ']') {
					continue
				} else {
					return false
				}
			}

		} else {
			return false
		}
	}
	return len(stack) == 0
}
