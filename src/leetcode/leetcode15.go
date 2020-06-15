package main

import (
	"fmt"
	s "sort"
)

/**
 存在重复的情况
     -1 -1  -1  2  2
	 -1 -1  0  1  1
     0 0 0


*/
func main() {
	fmt.Printf("%v\n", threeSum([]int{-1, 0, 1, 2, -1, -4})) // -4 -1 -1 0 1 2
	fmt.Printf("%v\n", threeSum([]int{0, 0, 0}))
	fmt.Printf("%v\n", threeSum([]int{-1, -1, -1, 2, 2}))
	fmt.Printf("%v", threeSum([]int{-1, -1, 0, 1, 1}))

}
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	s.Ints(nums)
	l := len(nums)
	for i := 0; i < l; {
		if i >= l-2 {
			break
		}
		for j, k := i+1, l-1; j < l || k > j; {
			temp := nums[j] + nums[k] + nums[i]
			if temp == 0 && j != k {
				temp := make([]int, 3)
				temp[0] = nums[i]
				temp[1] = nums[j]
				temp[2] = nums[k]
				result = append(result, temp)
				// 选取成功后  继续推进, 并且过滤掉有和重复的
				for {
					if j < l-1 && nums[j] == nums[j+1] {
						j++
					} else {
						break
					}
				}
				for {
					if k > 1 && nums[k] == nums[k-1] {
						k--
					} else {
						break
					}
				}
			}

			if temp > 0 {
				k--
			} else if temp < 0 {
				j++
			} else {
				k--
				j++
			}

			if j >= k {
				break
			}

		}

		for {
			if i < l-1 && nums[i] == nums[i+1] {
				i++
			} else {
				break
			}
		}
		i++

	}
	return result
}
