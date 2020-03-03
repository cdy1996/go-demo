package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", lengthOfLastWord("a  das "))

}

func lengthOfLastWord(s string) int {
	if len(s) == 0 || (len(s) == 1 && s[0] == ' ') {
		return 0
	}
	lastSum := 0
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if sum != 0 {
				lastSum = sum
			}
			sum = 0
			continue
		}
		sum++
		if sum != 0 {
			lastSum = sum
		}
	}
	return lastSum
}
