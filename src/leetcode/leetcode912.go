package main

import "fmt"

func main() {
	//fmt.Printf("%v \n", sortArray([]int{5, 2, 3, 1}))
	fmt.Printf("%v \n", sortArray1([]int{5, 1, 1, 2, 0, 0}))
}
func sortArray(nums []int) []int {
	sort(nums)
	return nums
}

func sort(nums []int) {
	if len(nums) < 2 {
		return
	}
	reference := nums[0]
	head, tail := 0, len(nums)-1
	i := 1
	referenceIndex := 0
	for head < tail {
		if nums[i] < reference {
			nums[i], nums[referenceIndex] = nums[referenceIndex], nums[i]
			head++
			i++
			referenceIndex++
		} else if nums[i] == reference {
			i++
			head++
		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		}
	}

	sort(nums[:referenceIndex])
	sort(nums[tail+1:])
}

func sortArray1(nums []int) []int {
	ret := []int{}
	arr := make(chan int, len(nums))
	Sort(nums, arr)
	for i := 0; i < len(nums); i++ {
		ret = append(ret, <-arr)
	}
	return ret
}

func Sort(arr []int, ch chan int) {
	defer close(ch)
	if len(arr) <= 1 {
		if len(arr) == 1 {
			ch <- arr[0]
		}
		return
	}
	mid := len(arr) / 2
	s1 := make(chan int, mid)
	s2 := make(chan int, len(arr)-mid)
	go Sort(arr[:mid], s1)
	go Sort(arr[mid:], s2)
	Merge(s1, s2, ch)
}

func update(s chan int, ch chan int, c *int, ok *bool) {
	ch <- *c
	*c, *ok = <-s
}

func Merge(s1, s2, ch chan int) {
	v1, ok1 := <-s1
	v2, ok2 := <-s2
	for ok1 && ok2 {
		if v1 < v2 {
			ch <- v1
			v1, ok1 = <-s1
		} else {
			ch <- v2
			v2, ok2 = <-s2
		}
	}
	for ok1 {
		ch <- v1
		v1, ok1 = <-s1
	}
	for ok2 {
		ch <- v2
		v2, ok2 = <-s2
	}
}
