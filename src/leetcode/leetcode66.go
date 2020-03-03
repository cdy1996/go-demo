package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", plusOne([]int{9, 9}))
	fmt.Printf("%v \n", plusOne([]int{1, 9}))

}

func plusOne(digits []int) []int {
	b := 0
	digits[len(digits)-1] = digits[len(digits)-1] + 1
	for i := len(digits) - 1; i >= 0; i-- {
		i2 := digits[i] + b
		if i2 >= 10 {
			b = 1
			digits[i] = i2 - 10
		} else {
			digits[i] = i2
			b = 0
		}
	}
	if b == 1 {
		ii := make([]int, len(digits)+1)
		ii[0] = 1
		for i := 0; i < len(digits); i++ {
			ii[i+1] = digits[i]
		}
		return ii
	}
	return digits

}
