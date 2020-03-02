package main

import "fmt"

// 字符串匹配 bm算法
// 大致逻辑
//		从模式串最后一位开始比较
//      字符移位: 坏字符串规则, 好后缀规则
// 小灰公众号
func main() {
	str := "bbbbababbbaabbba"
	pattern := "abb"
	index := boyerMoore(str, pattern)
	fmt.Printf("首次出现的位置 %d", index)

}

func boyerMoore(str, pattern string) int {
	strLength := len(str)
	patternLength := len(pattern)
	start := 0 //模式串起始位置
	for start <= strLength-patternLength {
		var i int
		for i = patternLength - 1; i >= 0; i-- {
			if str[start+i] != pattern[i] {
				break //发现坏字符
			}

		}
		if i < 0 {
			return start //匹配成功,返回第一次匹配的下标位置
		}
		// 寻找坏字符在模式串中的对应
		charIndex := findCharacter(pattern, str[start+i:start+i+1], i)

		if charIndex >= 0 { //如果存在坏值, 就把模式串的坏字符位置与主串的位置对应
			start += i - charIndex
		} else {
			start += i + 1 //如果不存在坏值, 就把模式串移到主串坏字符串的下一位
		}
	}
	return -1

}

// 在模式串中, 查找index下标之前的字符是否和坏字符串匹配
func findCharacter(pattern, bad string, index int) int {
	for i := index - 1; i >= 0; i-- {
		if pattern[i:i+1] == bad {
			return i

		}

	}
	return -1
}
