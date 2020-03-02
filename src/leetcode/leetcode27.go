package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//nums:=[]int{2,3}
	i := removeElement(nums, 2)
	for j := 0; j < i; j++ {
		fmt.Printf("%v,", nums[j])
	}
}
func removeElement(nums []int, val int) int {
	index := -1
	num := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			num++
			if index == -1 { // 如果index还没赋值, 那么将当前第一个重复的值的下标赋值给index
				index = i
			}
		} else {
			//前面没有重复的就不用赋值了
			if index != -1 { // 存在重复的值
				nums[index] = nums[i]
				index++ //移至下一位
			}
		}

	}
	return len(nums) - num
}
