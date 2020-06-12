package main

import (
	"fmt"
	s "sort"
)

/**
未完成
*/
func main() {
	fmt.Printf("%v", threeSum([]int{-1, 0, 1, 2, -1, -4}))

}
func threeSum(nums []int) [][]int {
	maps := make(map[int][]int)
	result := make([][]int, 0)
	s.Ints(nums)
	for i, num := range nums {
		v, b := maps[num]
		if !b {
			maps[num] = []int{i}
		} else {
			maps[num] = append(v, i)
		}
	}
	l := len(nums)
	for i, num := range nums {
		if l-i == 1 {
			break
		}
		for j := i + 1; j < l; j++ {
			b, v := maps[-num-nums[j]]
			if v {
				for i2 := range b {
					if i < j && j < b[i2] {
						temp := make([]int, 3)
						temp[0] = nums[i]
						temp[1] = nums[j]
						temp[2] = nums[b[i2]]

						result = append(result, temp)
						break
					}
				}

			}
		}
	}
	return result
}
