package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v", romanToInt("MMMM"))
}

func romanToInt(s string) int {
	maps := make(map[string]int)
	maps["I"] = 1
	maps["V"] = 5
	maps["X"] = 10
	maps["L"] = 50
	maps["C"] = 100
	maps["D"] = 500
	maps["M"] = 1000
	maps["IV"] = 4
	maps["IX"] = 9
	maps["XL"] = 40
	maps["XC"] = 90
	maps["CD"] = 400
	maps["CM"] = 900
	r := []rune(s)
	sum := 0
	i2 := len(r)
	for i := 0; i < i2; i++ {
		if r[i] == 'I' && i+1 < i2 && (r[i+1] == 'V' || r[i+1] == 'X') {
			sum += maps[string(r[i])+string(r[i+1])]
			i++
		} else if r[i] == 'X' && i+1 < i2 && (r[i+1] == 'L' || r[i+1] == 'C') {
			sum += maps[string(r[i])+string(r[i+1])]
			i++
		} else if r[i] == 'C' && i+1 < i2 && (r[i+1] == 'D' || r[i+1] == 'M') {
			sum += maps[string(r[i])+string(r[i+1])]
			i++
		} else {
			sum += maps[string(r[i])]
		}
	}
	return sum

}
