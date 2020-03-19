package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", maxProfit([]int{}))
	fmt.Printf("%v \n", maxProfit([]int{7}))
	fmt.Printf("%v \n", maxProfit([]int{7, 6}))
	fmt.Printf("%v \n", maxProfit([]int{7, 6, 4, 3, 1}))
}

//最大利益 自底向上
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	if len(prices) == 1 {
		return 0
	}

	maxOut := prices[len(prices)-1]
	max := maxOut - prices[len(prices)-2]
	if len(prices) == 2 {
		if max < 0 {
			return 0
		}
		return max
	}

	for i := len(prices) - 3; i >= 0; i-- {
		if prices[i+1] > maxOut {
			maxOut = prices[i+1]
		}
		// 如果 i后面最大的 - 第i天   比现在最大的大就是最大利润
		if maxOut-prices[i] > max {
			max = maxOut - prices[i]
		}
	}
	if max < 0 {
		return 0
	}
	return max
}
