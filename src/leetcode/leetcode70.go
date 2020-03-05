package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", climbStairs(44))

}
func climbStairsOld(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)

}
func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	arr := make([]int, n+1)
	arr[1] = 1
	arr[2] = 2
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
