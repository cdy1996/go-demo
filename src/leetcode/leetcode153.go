package main

import "fmt"

func main() {
	var ints = []int{4, 5, 6, 7, 0, 1, 2}

	cur := 0
	next := 1
	for next <= len(ints) {
		if ints[cur] > ints[next] {
			break
		}

		cur++
		next++
	}
	fmt.Printf("最小值为 %v", ints[next])

}
