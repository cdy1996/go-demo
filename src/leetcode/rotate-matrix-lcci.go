package main

import "fmt"

func main() {
	//var result [][]int
	//var currentLevel []int
	//currentLevel = append(currentLevel, 1)
	//
	//matrix:=make([][]int, 3)
	//matrix[0]=[]int{1,2,3}
	//matrix[1]=[]int{4,5,6}
	//matrix[2]=[]int{7,8,9}
	//rotate(matrix)

	matrix := make([][]int, 4)
	matrix[0] = []int{5, 1, 9, 11}
	matrix[1] = []int{2, 4, 8, 10}
	matrix[2] = []int{13, 3, 6, 7}
	matrix[3] = []int{15, 14, 12, 16}
	rotate(matrix)

	//matrix:=make([][]int, 2)
	//matrix[0]=[]int{1,2}
	//matrix[1]=[]int{3,4}
	//rotate(matrix)
	fmt.Printf("%v", matrix)

}
func rotate(matrix [][]int) {
	x := len(matrix)
	n := x - 1
	j := 0
	for {
		for i := j; i < n; i++ {
			matrix[i+j][n], matrix[n][n-i], matrix[n-i][j], matrix[j][i+j] =
				matrix[j][i+j], matrix[i+j][n], matrix[n][n-i], matrix[n-i][j]
		}
		x -= 2
		j++
		n--
		if x < 2 {
			break
		}
	}
}

func rotate1(matrix [][]int) {
	x := len(matrix)
	n := x - 1
	for i := 0; i < n; i++ {
		temp := matrix[n-i][0]
		matrix[n-i][0] = matrix[n][n-i]
		matrix[n][n-i] = matrix[i][n]
		matrix[i][n] = matrix[0][i]
		matrix[0][i] = temp
	}
	if x > 3 {
		var a11 [][]int
		for i := 1; i < n; i++ {
			a11 = append(a11, matrix[i][1:n])
		}
		rotate(a11)
	}

}
