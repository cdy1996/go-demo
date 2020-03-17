package main

import "fmt"

func main() {

	sum := twoSum1([]int{3, 2, 4}, 6)
	fmt.Printf("%v", sum)
}

func twoSum1(nums []int, target int) []int {
	map1 := make(map[int]int)
	for i, num := range nums {
		map1[num] = i
	}
	for i, num := range nums {
		if i2, ok := map1[target-num]; ok {
			if i == i2 {
				continue
			}
			return []int{i, i2}
		}
	}
	return nil
}
