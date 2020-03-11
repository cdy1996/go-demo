package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", singleNumber([]int{4, 1, 2, 1, 2}))

}

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	b := nums[0]
	for i := 1; i < len(nums); i++ {
		b = b ^ nums[i]
	}
	return b
}
