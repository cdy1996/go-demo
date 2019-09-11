package leetcode

import "fmt"

func main() {
	num:=10
	sum:=0
	for {
		sum=sum+num / 3 //已经换的总数
		i := num % 3 //多余的
		num = num/3 + i //换完之后的空瓶


		if num == 2 {
			sum++
			break
		}
		if num<=1 {
			break
		}

	}
	fmt.Println(sum)
}
