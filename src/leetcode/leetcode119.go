package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", getRow(3))
}

// 上台阶 自底向上
func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}
	ints := make([]int, 1)
	ints[0] = 1
	if rowIndex == 0 {
		return ints
	}
	ints = make([]int, 2)
	ints[0] = 1
	ints[1] = 1

	if rowIndex == 1 {
		return ints
	}
	for i := 2; i <= rowIndex; i++ {
		j := 1
		nextints := make([]int, i+1)
		nextints[0] = 1
		for j = 1; j < i; j++ {
			nextints[j] = ints[j-1] + ints[j]
		}
		nextints[i] = 1
		ints = append([]int{}, nextints...)
	}
	return ints
}
