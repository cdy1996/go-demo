package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Printf("%v \n", maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Printf("%v \n", maxSubArray([]int{-2, -2}))
}

func maxSubArray(nums []int) int {
	end := -1
	maxSum := math.MinInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		end = i
		sum = sum + nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 { //如果把前面都弄成 小于0, 可能前面的一些就是最大的了, 也可能后面还有更大的
			break
		}

	}
	if end == len(nums)-1 {
		return maxSum
	}
	sum1 := maxSubArray(nums[end+1:])
	if maxSum > sum1 {
		return maxSum
	} else {
		return sum1
	}
}
