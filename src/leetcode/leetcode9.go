package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v", isPalindrome1(121))
}

func isPalindrome1(x int) bool {
	x1 := x
	if x < 0 {
		return false
	}
	y := 0
	for x > 0 {
		y = y * 10
		i2 := x % 10
		x = x / 10
		y = y + i2
	}
	return x1 == y
}
