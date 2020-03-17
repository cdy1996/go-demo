package main

import "fmt"

func main() {

	//fmt.Printf("%v", twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Printf("%v", twoSum([]int{2, 3,4}, 6))
	fmt.Printf("%v", twoSum([]int{-1, 0}, -1))

}

func twoSum(numbers []int, target int) []int {
	len := len(numbers)
	if len == 0 || len == 1 {
		return []int{}
	}

	for i := len - 1; i > 0; i-- {

		for j := i - 1; j >= 0; j-- {
			if numbers[i]+numbers[j] == target {
				return []int{j + 1, i + 1}
			}
		}
	}
	return []int{}
}
