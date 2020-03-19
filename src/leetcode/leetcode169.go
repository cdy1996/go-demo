package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v \n", longestPalindrome("bb"))
}

// 出现最多的
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]

}
