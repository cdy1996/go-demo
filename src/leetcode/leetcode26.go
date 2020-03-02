package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	i := removeDuplicates(nums)
	for j := 0; j < i; j++ {
		fmt.Printf("%v,", nums[j])
	}
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	k := nums[0]
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == k {
			continue
		} else {
			nums[index] = nums[i]
			index++
			k = nums[i]

		}

	}
	return index
}
