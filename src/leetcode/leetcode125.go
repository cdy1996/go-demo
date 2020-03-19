package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%v \n", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("%v \n", isPalindrome("race a car"))

}

// 是否回文数字
func isPalindrome(s string) bool {
	r := []rune(s)
	l := len(r)

	for i, j := 0, l-1; i < l/2; {

		if !isNumOrLetter(r[i]) {
			i++
			continue
		}
		if !isNumOrLetter(r[j]) {
			j--
			continue
		}
		//fmt.Printf("%v->%v == %v->%v \n",i,r[i],j, r[j])
		if unicode.ToUpper(r[i]) != unicode.ToUpper(r[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
func isNumOrLetter(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)

}
