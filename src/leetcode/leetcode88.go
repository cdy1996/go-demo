package main

import (
	"fmt"
)

func main() {
	ints := make([]int, 1)
	merge(ints, 0, []int{1}, 1)
	fmt.Printf("%v \n", ints)

}
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}

	last := m + n - 1
	for {
		if m == 0 || nums1[m-1] <= nums2[n-1] {
			n--
			nums1[last] = nums2[n]
		} else {
			m--
			nums1[last] = nums1[m]
		}
		last--
		if n == 0 {
			break
		}
	}
}
