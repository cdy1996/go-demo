package main

import "fmt"

func main() {
	fmt.Printf("%v \n", findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Printf("%v \n", findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Printf("%v \n", findDuplicate([]int{1, 1, 2}))

}
func findDuplicate(nums []int) int {
	l := len(nums)
	start := 1
	end := l
	for {
		cnt := 0
		mid := (start + end) / 2
		for i := 0; i < l; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt == mid {
			if start == mid {
				start++
			} else {
				start = mid
			}
		} else if cnt > mid {
			if end == mid {
				end--
			} else {
				end = mid
			}
		} else if cnt < mid {
			if start == mid {
				start++
			} else {
				start = mid
			}
		}
		if start == end {
			return end
		}
	}
}
