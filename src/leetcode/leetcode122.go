package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v \n", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%v \n", maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Printf("%v \n", maxProfit([]int{}))
	fmt.Printf("%v \n", maxProfit([]int{1, 2}))
	fmt.Printf("%v \n", maxProfit([]int{1}))
	fmt.Printf("%v \n", maxProfit([]int{2, 2, 5}))
}

// 最大股票利益 贪心算法
func maxProfit(prices []int) int {

	sum := 0
	maxPos := -1
	minPos := -1

	len := len(prices)

	if len == 0 || len == 1 {
		return 0
	}

	for i := 0; i < len; i++ {
		// 凹点
		// 下一个没有了, 前一个比他大是凹点,  前一个没有了,后一个比他大也是凹点, 前后都比他大也是凹点
		if (i+1 == len && prices[i-1] >= prices[i]) || (i == 0 && prices[i+1] >= prices[i]) || (i != 0 && i+1 != len && prices[i+1] >= prices[i] && prices[i-1] >= prices[i]) {
			minPos = i
		}
		//凸点
		if (i+1 == len && prices[i-1] <= prices[i]) || (i == 0 && prices[i+1] <= prices[i]) || (i != 0 && i+1 != len && prices[i+1] <= prices[i] && prices[i-1] <= prices[i]) {
			maxPos = i
		}
		if minPos != -1 && maxPos != -1 && minPos < maxPos {
			sum += prices[maxPos] - prices[minPos]
			maxPos = -1
			minPos = -1
		}
	}
	return sum
}
