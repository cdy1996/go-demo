package main

import (
	"fmt"
)

/**
这题做复杂了, 最简单的用队列存t 就行了 , 每次获取队头看看是不是 < t-3000
*/
func main() {

	obj := Constructor()
	fmt.Printf("%v\n", obj.Ping(1))
	fmt.Printf("%v\n", obj.Ping(100))
	fmt.Printf("%v\n", obj.Ping(3000))
	fmt.Printf("%v\n", obj.Ping(3001))
	fmt.Printf("%v\n", obj.Ping(3002))
	fmt.Printf("%v\n", obj.Ping(3100))
	fmt.Printf("%v\n", obj.Ping(3101))

}

type RecentCounter struct {
	total     int
	array     []int
	round     int
	lastIndex int
}

func Constructor() RecentCounter {
	return RecentCounter{
		total:     0,
		array:     make([]int, 3002),
		round:     0,
		lastIndex: 0,
	}
}

func (this *RecentCounter) Ping(t int) int {
	// 获取下标
	round := t / 3001
	sub := round - this.round
	index := t % 3001
	// 如果 大于1  全部设置为0
	if sub > 1 {
		for i := 0; i < len(this.array); i++ {
			this.array[i] = 0
		}
		this.total = 0
	} else {
		// 如果 等于0 说明是同一个轮次
		// 如果 等于1 说明进入下一轮次, 也就是从 lastIndex -> index 的位置要全部设置为0
		for i := this.lastIndex + 1; i <= t; i++ {
			this.total = this.total - this.array[i%3001]
			this.array[i%3001] = 0
		}
	}
	this.array[index]++
	this.total++
	this.round = round
	this.lastIndex = t
	return this.total
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
