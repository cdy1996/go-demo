package main

import "fmt"

/*
数组 传递
*/
func main() {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The original array: %v\n", array1)
	fmt.Printf("The modified array: %v\n", array2)

	array3 := modifyArray2(&array1)
	fmt.Printf("The original array: %v\n", array1)
	fmt.Printf("The modified array: %v\n", array3)

	array4 := make([]string, 3)
	array4[0] = "a"
	array4[1] = "b"
	array4[2] = "c"
	array5 := modifyArray3(array4)
	fmt.Printf("The original array: %v\n", array4)
	fmt.Printf("The modified array: %v\n", array5)

}
func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}
func modifyArray2(a *[3]string) *[3]string {
	a[1] = "x"
	return a
}
func modifyArray3(a []string) []string {
	a[1] = "x"
	return a
}
