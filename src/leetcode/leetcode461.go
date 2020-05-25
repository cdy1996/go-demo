package main

import "fmt"

func main() {
	fmt.Printf("%v \n", hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	i := x ^ y
	j := 0
	for {
		if i&1 == 1 {
			j++
		}
		i = i >> 1
		if i == 0 {
			return j
		}

	}
	return j
}
