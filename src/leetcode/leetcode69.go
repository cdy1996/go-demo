package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v \n", mySqrt(9))
	fmt.Printf("%v \n", mySqrt(10))
	fmt.Printf("%v \n", mySqrt(0))
	fmt.Printf("%v \n", mySqrt(1))
	fmt.Printf("%v \n", mySqrt(100))
	fmt.Printf("%v \n", mySqrt(1024))

}
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	xx := x
	i := 0
	for xx > 0 {
		xx = xx / 10
		i++
	}

	i2 := i / 2
	i3 := i2 + 1
	if i == 1 {
		i2 = 1
		i3 = 1
	}
	start := int(math.Pow(10, float64(i2-1)))
	end := int(math.Pow(10, float64(i3)))

	for j := start; j < end; j++ {
		if j*j > x {
			return j - 1
		}
	}
	return 0
}
