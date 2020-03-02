package main

import "fmt"

// 字符串匹配 jmp算法
//https://blog.csdn.net/dark_cy/article/details/88698736
//https://www.cnblogs.com/zhangtianq/p/5839909.html
//https://www.cnblogs.com/dusf/p/kmp.html
func main() {
	str := "bbbbababbbaabbba"
	pattern := "abb"
	//kmp := KMP("BBC ABCDAB ABCDABCDABDE", "ABCDABD")
	kmp := KMP(str, pattern)
	fmt.Println()
	fmt.Println(kmp)
}

// next数组
func GetNext(next []int, t string) {
	var k = -1
	next[0] = -1
	for j := 0; j < len(next)-1; {

		if k == -1 || t[j] == t[k] {
			j++
			k++
			next[j] = k

			//if(t[j]==t[k]) { //当两个字符相同时，就跳过
			//	next[j] = next[k];
			//} else {
			//	next[j] = k;
			//}
		} else {
			k = next[k] //回退 k
		}
		for e := range next {
			fmt.Print(next[e])
			fmt.Print(",")
		}
		fmt.Println()
	}
}

// kmp算法
// s
// t 子串
func KMP(s, t string) int {
	var next = make([]int, len(t))
	GetNext(next, t)
	var i, j = 0, 0
	for i < len(s) && j < len(t) {
		if j == -1 || s[i] == t[j] {
			i++
			j++

		} else {
			j = next[j]

		}
	}
	if j >= len(t) {
		return i - len(t)
	} else {
		return -1
	}
}
