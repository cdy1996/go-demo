package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v \n", countAndSay(5))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := "1"
	for i := 2; i <= n; i++ {
		ss := ""  //临时字符串
		k := 0    //相同的个数
		t := s[0] // 用于判断是否相同的
		for j := 0; j < len(s); j++ {
			if t == s[j] {
				k++
			} else {
				ss = ss + strconv.Itoa(k) + string(s[j-1])
				t = s[j]
				k = 1
			}
		}
		if k != 0 { //用于最后一位的描述
			ss = ss + strconv.Itoa(k) + string(s[len(s)-1])
		}
		s = ss
	}
	return s
}
