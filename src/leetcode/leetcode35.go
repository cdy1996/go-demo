package main

import "fmt"

func main() {
	fmt.Printf("%v \n", searchInsert([]int{3}, 2))
}

func searchInsert(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return i
		}

	}
	return len(nums)
}
