package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", generate(0))
}

// 斐波那契数列
func generate(numRows int) [][]int {
	var ints = make([][]int, numRows)
	if numRows <= 0 {
		return ints
	}
	ints[0] = make([]int, 1)
	ints[0][0] = 1
	if numRows == 1 {
		return ints
	}
	ints[1] = make([]int, 2)
	ints[1][0] = 1
	ints[1][1] = 1

	if numRows == 2 {
		return ints
	}

	for i := 2; i < numRows; i++ {
		ints[i] = make([]int, i+1)
		ints[i][0] = 1
		j := 1
		for j = 1; j < i; j++ {
			ints[i][j] = ints[i-1][j-1] + ints[i-1][j]
		}
		ints[i][j] = 1
	}
	return ints
}
