package main

import "fmt"

func main() {
	var ints = []int{2, 3, 1, 1, 4}
	//ints := make([]int, 5)

	fmt.Printf("是否能到底 %v", toEnd(ints))

}

func toEnd(ints []int) bool {
	current := 0
	return jump(current, ints) == 0
}

func jump(current int, ints []int) int {
	if current == len(ints)-1 { // 已经到达尾部 则成功
		return 0
	}
	if ints[current] == 0 { //如果当前为0 则不需要走
		return -1
	}
	for i := 1; i <= ints[current]; i++ { // 循环多种走法
		if current+i == len(ints)-1 { // 如果下一步到尾 则不需要走
			return 0
		}
		if ints[current+i] == 0 { //如果下一个为0 则不需要走
			continue
		}
		i2 := jump(current+i, ints)
		if i2 == 0 {
			return 0
		}
	}
	return -1
}
