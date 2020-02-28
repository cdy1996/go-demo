package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v", reverse(1534236469))
}

func reverse(x int) int {
	i := 1
	if x > 0 {
		i = 1
	} else {
		x = x * -1
		i = -1
	}
	y := 0

	for x > 0 {
		y = y * 10
		i2 := x % 10
		x = x / 10
		y = y + i2
	}
	i2 := y * i
	if i > 0 {
		if i2 > math.MaxInt32 {
			return 0
		}
	} else {
		if i2 < math.MinInt32 {
			return 0
		}
	}
	return i2
}
