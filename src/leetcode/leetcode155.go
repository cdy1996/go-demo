package main

import (
	"fmt"
	"math"
)

func main() {

	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Printf("%v \n", stack.GetMin())
	stack.Pop()
	fmt.Printf("%v \n", stack.Top())
	fmt.Printf("%v \n", stack.GetMin())

}

// 最小栈
type MinStack struct {
	Min   int
	Array []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Min:   math.MaxInt32,
		Array: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if this.Min > x {
		this.Min = x
	}
	this.Array = append(this.Array, x)
}

func (this *MinStack) Pop() {
	len := len(this.Array)
	if len == 0 {
		return
	}
	if len == 1 {
		this.Min = math.MaxInt32
		this.Array = make([]int, 0)
		return
	}
	if this.Array[len-1] == this.Min {
		min := math.MaxInt32
		for i := 0; i < len-1; i++ {
			if min > this.Array[i] {
				min = this.Array[i]
			}
		}
		this.Min = min
	}
	this.Array = this.Array[0 : len-1]

}

func (this *MinStack) Top() int {
	if len(this.Array) == 0 {
		panic("have no num")
	}

	return this.Array[len(this.Array)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Array) == 0 {
		panic("have no num")
	}

	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
